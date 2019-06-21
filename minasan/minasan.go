package minasan

import (
	"fmt"
	"io"
	"strings"

	"net/smtp"
	"net/textproto"

	"github.com/factorysh/minasan/gitlab"
	"github.com/factorysh/minasan/metrics"
	"github.com/flashmob/go-guerrilla/mail"
	log "github.com/sirupsen/logrus"
)

// Minasan sends emails from Gitlab groups to a SMTPOut
type Minasan struct {
	Client       *gitlab.Client
	SMTPOut      string
	Bcc          string
	SenderDomain string
}

// Targets return mails, group, project from a mail name
func (m *Minasan) Targets(mailName, lastChanceMail string) ([]string, string, string, error) {
	blob := strings.Split(mailName, ".")
	if len(blob) != 2 {
		return nil, "", "", fmt.Errorf("Bad mail name, can't guess group and project : %s", mailName)
	}
	group := blob[0]
	project := blob[1]
	targets, err := m.Client.MailsFromGroupProject(group, project, lastChanceMail)
	if err != nil {
		return nil, "", "", err
	}
	return targets, group, project, nil
}

func writeStuff(w io.Writer, blobs ...string) error {
	for _, blob := range blobs {
		_, err := w.Write([]byte(blob))
		if err != nil {
			return err
		}
	}
	return nil
}

// BroadcastMail to all targets
func (m *Minasan) BroadcastMail(mails []string, envelope *mail.Envelope,
	header textproto.MIMEHeader) error {
	// https://github.com/golang/go/wiki/SendingMail
	if m.Bcc != "" {
		mails = append(mails, m.Bcc)
	}
	for _, mail := range mails {
		if mail == "" {
			log.Warning("Empty mail, it's a real bug, but I don't want to crash. FIXME")
			continue
		}
		c, err := smtp.Dial(m.SMTPOut)
		if err != nil {
			log.WithField("smtpout", m.SMTPOut).WithError(err).Error("Can't dial SMTP")
			return err
		}
		defer c.Close()
		from := envelope.MailFrom.String()
		from = fmt.Sprintf("minasan+%s@%s", strings.Split(from, "@")[0], m.SenderDomain)
		c.Mail(from)
		c.Rcpt(mail)
		wc, err := c.Data()
		if err != nil {
			log.WithFields(log.Fields{
				"mail": envelope.MailFrom.String(),
				"rcpt": mail,
			}).Error(err)
			return err
		}
		defer wc.Close()
		for key, values := range header {
			for _, value := range values {
				err := writeStuff(wc, key, ": ", value, "\n")
				if err != nil {
					log.WithError(err).WithField(key, value).Error("Can't write header")
					return err
				}
			}
		}
		_, err = io.Copy(wc, envelope.NewReader())
		if err != nil {
			log.WithError(err).Error("Can't copy envelope")
			return err
		}
		metrics.MailSentCounter.Inc()
	}
	log.WithFields(log.Fields{
		"mails":   strings.Join(mails, ", "),
		"subject": envelope.Subject,
		"from":    envelope.MailFrom.String(),
	}).Info("BroadcastMail")
	return nil
}
