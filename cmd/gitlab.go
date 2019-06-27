package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/factorysh/minasan/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gitlabCmd)
	gitlabCmd.AddCommand(pingCmd)
}

var gitlabCmd = &cobra.Command{
	Use:   "gitlab [group] [project]",
	Short: "Ask gitlab wich mails are linked to a specific project",
	Args:  cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := gitlab.NewClientWithGitlabPrivateToken(nil,
			viper.GetString("gitlab_domain"),
			viper.GetString("gitlab_private_token"),
			5*time.Minute, "/tmp/minasan.db")
		if err != nil {
			return err
		}
		mails, err := client.MailsFromGroupProject(args[0], args[1])
		if err != nil {
			return err
		}
		for _, mail := range mails {
			fmt.Println(mail)
		}
		return nil
	},
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping your gitlab",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := gitlab.NewClientWithGitlabPrivateToken(nil,
			viper.GetString("gitlab_domain"),
			viper.GetString("gitlab_private_token"),
			5*time.Minute, "/tmp/minasan.db")
		if err != nil {
			return err
		}
		name, err := client.Ping()
		if err != nil {
			return err
		}
		fmt.Println(name)
		return nil
	},
}
