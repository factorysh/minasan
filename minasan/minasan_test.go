package minasan

import (
	"bytes"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flashmob/go-guerrilla/mail"
)

// You need `make mailhog` before running this test
func TestMail(t *testing.T) {
	m := &Minasan{
		SMTPOut: "0.0.0.0:1025",
	}

	addr, err := mail.NewAddress("Sender <sender@example.com>")
	assert.NoError(t, err)
	envelope := mail.NewEnvelope("127.0.0.1", 42)
	envelope.Data = *bytes.NewBufferString(`
Hello World
`)
	envelope.Subject = "Big test"
	envelope.MailFrom = addr
	header := make(textproto.MIMEHeader)
	m.BroadcastMail([]string{
		"pim@example.com",
		"pam@example.com",
		"poum@example.com"},
		envelope, header)
}
