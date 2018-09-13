package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "minasan",
	Short: "Send mail to gitlab projects",
	Long: `Environement variables:

  SMTP_DOMAIN: example.com
  GITLAB_DOMAIN: gitlab.example.com
  GITLAB_PRIVATE_TOKEN: oppoipoioipipoipo
  SMTP_OUT: localhost:25
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
