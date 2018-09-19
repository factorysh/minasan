package cmd

import (
	"fmt"

	"github.com/factorysh/minasan/gitlab"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gitlabCmd)
}

var gitlabCmd = &cobra.Command{
	Use:   "gitlab [group] [project]",
	Short: "Ask gitlab wich mails are linked to a specific project",
	Args:  cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		client := gitlab.NewClientFromEnv(nil)
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
