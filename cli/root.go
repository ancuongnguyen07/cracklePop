package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cracklepop",
	Short: "Cracklepop CLI tool",
	Long:  "Cracklepop CLI tool",
	Example: `
	cracklepop print --start 1 --end 100
	`,
	Version: "v0",
	Run:     root,
}

func root(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func Execute() error {
	return rootCmd.Execute()
}
