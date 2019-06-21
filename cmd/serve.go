package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/viper"

	"github.com/factorysh/minasan/metrics"
	"github.com/factorysh/minasan/processor"
	guerrilla "github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	smtpDomain    string
	smtpIn        string
	smtpOut       string
	metricsAdress string
	bcc           string
	senderDomain  string
	returnPath    string
)

func init() {
	pf := serveCmd.PersistentFlags()
	pf.StringVarP(&smtpDomain, "smtp_domain", "d", "gitlab.example.com", "SMTP domain")
	pf.StringVarP(&smtpIn, "smtp_in", "i", "127.0.0.1:2525", "SMTP input service")
	pf.StringVarP(&smtpOut, "smtp_out", "o", "127.0.0.1:25", "SMTP relay")
	pf.StringVarP(&metricsAdress, "metrics_address", "H", "127.0.0.1:8125", "Prometheus probe listening address")
	pf.StringVarP(&bcc, "bcc", "b", "", "Blind carbon copy")
	pf.StringVarP(&senderDomain, "sender_domain", "s", "example.com", "Sender domain, for the smtp relay, admin@{sender_domain} will be used")
	pf.StringVarP(&returnPath, "return_path", "r", "admin@example.com", "Return path")
	viper.BindPFlag("smtp_domain", serveCmd.PersistentFlags().Lookup("smtp_domain"))
	viper.BindPFlag("smtp_in", serveCmd.PersistentFlags().Lookup("smtp_in"))
	viper.BindPFlag("smtp_out", serveCmd.PersistentFlags().Lookup("smtp_out"))
	viper.BindPFlag("metrics_address", serveCmd.PersistentFlags().Lookup("metrics_address"))
	viper.BindPFlag("bcc", serveCmd.PersistentFlags().Lookup("bcc"))
	viper.BindPFlag("sender_domain", serveCmd.PersistentFlags().Lookup("sender_domain"))
	viper.BindPFlag("return_path", serveCmd.PersistentFlags().Lookup("return_path"))
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Listen as a SMTP server",

	RunE: func(cmd *cobra.Command, args []string) error {
		ma := viper.GetString("metrics_address")
		if ma != "" {
			go metrics.ListenAndServe(ma)
		} else {
			log.Info("No prometheus probe")
		}
		d := guerrilla.Daemon{}
		d.SetConfig(guerrilla.AppConfig{
			AllowedHosts: []string{viper.GetString("smtp_domain")},
			BackendConfig: backends.BackendConfig{
				"save_process":         "HeadersParser|Debugger|Minasan",
				"validate_process":     "Minasan",
				"gitlab_domain":        viper.GetString("gitlab_domain"),
				"gitlab_private_token": viper.GetString("gitlab_private_token"),
				"last_chance_mail":     viper.GetString("last_chance_mail"),
				"smtp_out":             viper.GetString("smtp_out"),
				"bcc":                  viper.GetString("bcc"),
				"sender_domain":        viper.GetString("sender_domain"),
				"return_path":          viper.GetString("return_path"),
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
