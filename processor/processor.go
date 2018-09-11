package processor

import (
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
)

type myFooConfig struct {
	SomeOption string `json:"maildir_path"`
}

// The MyFoo decorator [enter what it does]
var MinasanProcessor = func() backends.Decorator {
	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				if task == backends.TaskValidateRcpt {

					// if you want your processor to validate recipents,
					// validate recipient by checking
					// the last item added to `e.RcptTo` slice
					// if error, then return something like this:
					/* return backends.NewResult(
					   response.Canned.FailNoSenderDataCmd),
					   backends.NoSuchUser
					*/
					// if no error:
					return p.Process(e, task)
				} else if task == backends.TaskSaveMail {

					// if you want your processor to do some processing after
					// receiving the email, continue here.
					// if want to stop processing, return
					// errors.New("Something went wrong")
					// return backends.NewBackendResult(fmt.Sprintf("554 Error: %s", err)), err
					// call the next processor in the chain
					return p.Process(e, task)
				}
				return p.Process(e, task)
			},
		)
	}
}
