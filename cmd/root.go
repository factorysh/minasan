package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gitlabDomain       string
	lastChanceMail     string
	gitlabPrivateToken string
	configFile         string
	debug              bool
	cachePath          string
	cacheExpiration    time.Duration
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
		if debug {
			log.SetLevel(log.DebugLevel)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&gitlabDomain, "gitlab_domain", "g", "gitlab.example.com", "Gitlab domain")
	pf.StringVarP(&lastChanceMail, "last_chance_mail", "m", "test@example.com", "Last chance mail")
	pf.StringVarP(&cachePath, "cache_path", "p", "/var/lib/minasan/minasan.db", "Gitlab cache path and name for the db")
	pf.DurationVarP(&cacheExpiration, "cache_expiration", "e", 5*time.Minute, "Expiration time for the gitlab cache")
	pf.StringVarP(&gitlabPrivateToken, "gitlab_private_token", "t", "", "Gitlab private token")
	pf.StringVarP(&configFile, "config", "c", "", "Config file")
	pf.BoolVarP(&debug, "verbose", "V", false, "More verbose")
	viper.BindPFlag("gitlab_domain", rootCmd.PersistentFlags().Lookup("gitlab_domain"))
	viper.BindPFlag("last_chance_mail", rootCmd.PersistentFlags().Lookup("last_chance_mail"))
	viper.BindPFlag("cache_path", rootCmd.PersistentFlags().Lookup("cache_path"))
	viper.BindPFlag("cache_expiration", rootCmd.PersistentFlags().Lookup("cache_expiration"))
	viper.BindPFlag("gitlab_private_token", rootCmd.PersistentFlags().Lookup("gitlab_private_token"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
