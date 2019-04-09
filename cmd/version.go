package cmd

import (
	"fmt"

	"github.com/factorysh/minasan/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of minasan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version())
	},
}
