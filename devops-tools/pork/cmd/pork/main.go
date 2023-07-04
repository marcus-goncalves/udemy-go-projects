package main

import (
	"github.com/spf13/cobra"
	pork "udemy-projects.com/devops-tools/pork/vendor"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "pork",
		Short: "project forking tool for github",
	}

	rootCmd.AddCommand(pork.SearchCmd)
	rootCmd.AddCommand(pork.DocsCmd)
	rootCmd.AddCommand(pork.CloneCmd)
	rootCmd.AddCommand(pork.ForkCmd)
}
