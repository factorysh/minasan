package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/viper"

	guerrilla "github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.bearstech.com/factory/minasan/processor"
)

var (
	smtpDomain string
	smtpIn     string
	smtpOut    string
)

func init() {
	pf := serveCmd.PersistentFlags()
	pf.StringVarP(&smtpDomain, "smtp_domain", "d", "gitlab.example.com", "SMTP domain")
	pf.StringVarP(&smtpIn, "smtp_in", "i", "127.0.0.1:2525", "SMTP input service")
	pf.StringVarP(&smtpOut, "smtp_out", "o", "127.0.0.1:25", "SMTP relay")
	viper.BindPFlag("smtp_domain", serveCmd.PersistentFlags().Lookup("smtp_domain"))
	viper.BindPFlag("smtp_in", serveCmd.PersistentFlags().Lookup("smtp_in"))
	viper.BindPFlag("smtp_out", serveCmd.PersistentFlags().Lookup("smtp_out"))
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Listen as a SMTP server",

	RunE: func(cmd *cobra.Command, args []string) error {
		d := guerrilla.Daemon{}
		d.SetConfig(guerrilla.AppConfig{
			AllowedHosts: []string{viper.GetString("smtp_domain")},
			BackendConfig: backends.BackendConfig{
				"save_process":         "HeadersParser|Debugger|Minasan",
				"validate_process":     "Minasan",
				"gitlab_domain":        viper.GetString("gitlab_domain"),
				"gitlab_private_token": viper.GetString("gitlab_private_token"),
				"smtp_out":             viper.GetString("smtp_out"),
			},
			Servers: []guerrilla.ServerConfig{
				guerrilla.ServerConfig{
					ListenInterface: viper.GetString("smtp_in"),
					IsEnabled:       true,
				},
			},
		})

		d.AddProcessor("Minasan", processor.MinasanProcessor)
		err := d.Start()

		if err == nil {
			log.WithFields(log.Fields{
				"gitlab_domain": viper.GetString("gitlab_domain"),
				"smtp_domain":   viper.GetString("smtp_domain"),
				"smtp_in":       viper.GetString("smtp_in"),
				"smtp_out":      viper.GetString("smtp_out"),
			}).Info("Server started")
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
