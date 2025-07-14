package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple Todo application using command line interface.",
	Long: `ToDo is a demonstration create and manage your tasks on command-line interface which developed with the Cobra library in Go.

It supports basic commands and flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to ToDo app with CLI! Use 'todo help' for more information.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}
