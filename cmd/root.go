package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gitlabDomain       string
	gitlabPrivateToken string
)

var rootCmd = &cobra.Command{
	Use:   "minasan",
	Short: "Send mail to gitlab projects",
}

func initConfig() {
	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize(initConfig)
	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&gitlabDomain, "gitlab_domain", "g", "gitlab.example.com", "Gitlab domain")
	pf.StringVarP(&gitlabPrivateToken, "gitlab_private_token", "t", "", "Gitlab private token")
	viper.BindPFlag("gitlab_domain", rootCmd.PersistentFlags().Lookup("gitlab_domain"))
	viper.BindPFlag("gitlab_private_token", rootCmd.PersistentFlags().Lookup("gitlab_private_token"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
