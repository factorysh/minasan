package dkim

import (
	"fmt"

	dkim "github.com/emersion/go-dkim"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
	"github.com/flashmob/go-guerrilla/response"
)

type myDkimConfig struct {
	SignatureIsMandatory bool
}

var DKIMProcessor = func() backends.Decorator {
	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				if task == backends.TaskSaveMail {
					verifications, err := dkim.Verify(e.NewReader())
					if err != nil {
						return backends.NewResult(
							response.Canned.FailReadErrorDataCmd,
						), backends.NoSuchUser
					}
					fmt.Println(verifications)
				}
				return p.Process(e, task)
			},
		)
	}
}
