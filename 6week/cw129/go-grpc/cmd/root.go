package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Long: `Sample Cobra for run gRPC Server.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
