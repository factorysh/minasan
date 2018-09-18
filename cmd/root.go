package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gitlabDomain       string
	gitlabPrivateToken string
	configFile         string
)

var rootCmd = &cobra.Command{
	Use:   "minasan",
	Short: "Send mail to gitlab projects",
}

func initConfig() {
	viper.AutomaticEnv()
	cfg := viper.GetString("config")
	if cfg != "" {
		abs, err := filepath.Abs(cfg)
		if err != nil {
			panic(err)
		}
		base := filepath.Base(abs)
		path := filepath.Dir(abs)
		viper.SetConfigName(strings.Split(base, ".")[0])
		viper.AddConfigPath(path)
		err = viper.ReadInConfig()
		log.WithFields(log.Fields{
			"path": path,
			"base": base,
		}).Info("Reading config file")
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&gitlabDomain, "gitlab_domain", "g", "gitlab.example.com", "Gitlab domain")
	pf.StringVarP(&gitlabPrivateToken, "gitlab_private_token", "t", "", "Gitlab private token")
	pf.StringVarP(&configFile, "config", "c", "", "Config file")
	viper.BindPFlag("gitlab_domain", rootCmd.PersistentFlags().Lookup("gitlab_domain"))
	viper.BindPFlag("gitlab_private_token", rootCmd.PersistentFlags().Lookup("gitlab_private_token"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
