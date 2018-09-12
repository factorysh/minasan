package minasan

import (
	"bytes"
	"testing"

	"github.com/flashmob/go-guerrilla/mail"
)

// You need `make mailhog` before running this test
func TestMail(t *testing.T) {
	m := &Minasan{
		SMTPOut: "127.0.0.1:1025",
	}

	envelope := mail.NewEnvelope("sender@example.com", 42)
	envelope.Data = *bytes.NewBufferString(`From: sender.example.com
Subject: Big test

Hello World
`)
	m.BroadcastMail([]string{"pim@example.com", "pam@example.com", "poum@example.com"}, envelope)
}
