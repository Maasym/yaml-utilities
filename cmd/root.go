package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yaml-linter",
	Short: "YAML Linter CLI",
	Long:  "YAML Linter is a CLI tool to validate and lint YAML files.",
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior when no subcommands are provided
		fmt.Println("Welcome to YAML Linter! Use --help for available commands.")
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
