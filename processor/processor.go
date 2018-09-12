package processor

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
	"github.com/flashmob/go-guerrilla/response"
	log "github.com/sirupsen/logrus"
	"gitlab.bearstech.com/factory/minasan/gitlab"
	minasan_ "gitlab.bearstech.com/factory/minasan/minasan"
)

type myMinasanConfig struct {
	GitlabDomain       string `json:"gitlab_domain"`
	GitlabPrivateToken string `json:"gitlab_private_token,omitempty"`
	SMTPOut            string `json:"smtp_out,omitempty"`
}

// The MyFoo decorator [enter what it does]
var MinasanProcessor = func() backends.Decorator {
	config := &myMinasanConfig{}
	minasan := &minasan_.Minasan{}
	// our initFunc will load the config.
	initFunc := backends.InitializeWith(func(backendConfig backends.BackendConfig) error {
		configType := backends.BaseConfig(&myMinasanConfig{})
		bcfg, err := backends.Svc.ExtractConfig(backendConfig, configType)
		if err != nil {
			return err
		}
		config = bcfg.(*myMinasanConfig)
		minasan.SMTPOut = config.SMTPOut
		minasan.Client = gitlab.NewClientWithGitlabPrivateToken(nil, config.GitlabDomain, config.GitlabPrivateToken)
		return nil
	})
	// register our initializer
	backends.Svc.AddInitializer(initFunc)
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
					spew.Dump(e, e.Header, config)
					targets, err := minasan.Targets(e.RcptTo[0].User)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err,
							"rcpt":  e.RcptTo,
						}).Error("Minia Validate")
						return backends.NewResult(
								response.Canned.FailNoSenderDataCmd),
							backends.NoSuchUser
					}
					e.Values["targets"] = targets
					// if no error:
					return p.Process(e, task)
				} else if task == backends.TaskSaveMail {

					// if you want your processor to do some processing after
					// receiving the email, continue here.
					// if want to stop processing, return
					// errors.New("Something went wrong")
					// return backends.NewBackendResult(fmt.Sprintf("554 Error: %s", err)), err
					// call the next processor in the chain
					targets := e.Values["targets"].([]string)
					err := minasan.BroadcastMail(targets, e)
					if err != nil {
						log.WithFields(log.Fields{
							"error":   err,
							"targets": targets,
						}).Error("Minia SaveMail")
						return backends.NewResult(
								response.Canned.FailNoSenderDataCmd),
							backends.RcptError(err)
					}
					return p.Process(e, task)
				}
				return p.Process(e, task)
			},
		)
	}
}
