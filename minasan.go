package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	"gitlab.bearstech.com/factory/minasan/gitlab"
	"gitlab.bearstech.com/factory/minasan/processor"
)

func doSmtpd() {
	d := guerrilla.Daemon{}
	d.SetConfig(guerrilla.AppConfig{
		AllowedHosts: []string{"example.com"},
		BackendConfig: backends.BackendConfig{
			"save_process":         "HeadersParser|Debugger|Minasan",
			"validate_process":     "Minasan",
			"gitlab_domain":        os.Getenv("GITLAB_DOMAIN"),
			"gitlab_private_token": os.Getenv("GITLAB_PRIVATE_TOKEN"),
			"smtp_out":             os.Getenv("SMTP_OUT"),
		},
	})
	d.AddProcessor("Minasan", processor.MinasanProcessor)
	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	} else {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		// sig is a ^C, handle it
		fmt.Println(sig)
		return
	}

}

func doGitlab() {
	client := gitlab.NewClientFromEnv(nil)
	mails, err := client.MailsFromGroupProject(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	for _, mail := range mails {
		fmt.Println(mail)
	}
}

func main() {
	if len(os.Args) > 1 {
		doGitlab()
	} else {
		doSmtpd()
	}
}
