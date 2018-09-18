package minasan

import (
	"fmt"
	"io"
	"strings"

	"net/smtp"
	"net/textproto"

	"github.com/flashmob/go-guerrilla/mail"
	log "github.com/sirupsen/logrus"
	"gitlab.bearstech.com/factory/minasan/gitlab"
)

type Minasan struct {
	Client  *gitlab.Client
	SMTPOut string
}

func (m *Minasan) Targets(mailName string) ([]string, string, string, error) {
	blob := strings.Split(mailName, ".")
	if len(blob) != 2 {
		return nil, "", "", fmt.Errorf("Bad mail name, can't guess group and project : %s", mailName)
	}
	group := blob[0]
	project := blob[1]
	targets, err := m.Client.MailsFromGroupProject(group, project)
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

func (m *Minasan) BroadcastMail(mails []string, envelope *mail.Envelope, header textproto.MIMEHeader) error {
	// https://github.com/golang/go/wiki/SendingMail
	for _, mail := range mails {
		c, err := smtp.Dial(m.SMTPOut)
		if err != nil {
			log.WithFields(log.Fields{
				"smtpout": m.SMTPOut,
				"error":   err,
			})
			return err
		}
		defer c.Close()
		c.Mail(envelope.MailFrom.String())
		c.Rcpt(mail)
		wc, err := c.Data()
		if err != nil {
			return err
		}
		defer wc.Close()
		/*
			_, err = wc.Write(envelope.Data.Bytes())
			if err != nil {
				return err
			}
		*/
		for k, v := range header {
			envelope.Header[k] = v
		}
		for key, values := range envelope.Header {
			for _, value := range values {
				err := writeStuff(wc, key, ": ", value, "\n")
				if err != nil {
					return err
				}
			}
		}
		err = writeStuff(wc, "Subject: ", envelope.Subject, "\n\n")
		if err != nil {
			return err
		}
		_, err = io.Copy(wc, envelope.NewReader())
		if err != nil {
			return err
		}
	}
	log.WithFields(log.Fields{
		"mails":   strings.Join(mails, ", "),
		"subject": envelope.Subject,
		"from":    envelope.MailFrom.String(),
	}).Info("BroadcastMail")
	return nil
}
