package cmd

import (
	"fmt"
	"os"
	"os/signal"

	guerrilla "github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/spf13/cobra"
	"gitlab.bearstech.com/factory/minasan/processor"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Listen as a SMTP server",

	RunE: func(cmd *cobra.Command, args []string) error {
		d := guerrilla.Daemon{}
		d.SetConfig(guerrilla.AppConfig{
			AllowedHosts: []string{os.Getenv("SMTP_DOMAIN")},
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
			return err
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for sig := range c {
			// sig is a ^C, handle it
			fmt.Println(sig)
			return nil
		}
		return nil
	},
}
