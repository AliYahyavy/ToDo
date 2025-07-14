package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"todo/handlers"
)

var (
	IsDone bool

	addCmd = &cobra.Command{
		Use:  "add <task description>",
		Long: `The 'add' command add a new task.It take's an description as parameter'`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			affcted := handlers.Add(args[0], IsDone)
			if affcted {
				fmt.Printf("Task '%s' added to todo list successfully.\n", args[0])
			}
		},
	}
)

func init() {
	addCmd.Flags().BoolVarP(&IsDone, "done", "d", false, "add finished task to todo list.")
}
