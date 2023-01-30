package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "convert",
	Short: "This CLI tool is a simple file converter",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// add subcommands
	rootCmd.AddCommand(CArray())
}
