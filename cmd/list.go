package cmd

import (
	"github.com/spf13/cobra"
	"todo/handlers"
	"todo/utils"
)

var (
	showAll    bool
	listCount  int
	descending bool

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "return a list of all of the uncompleted tasks",
		Long:  "list command return a list of all of the uncompleted tasks, there some for features like sort, list completed tasks and list arbitrary line of tasks.",
		Run: func(cmd *cobra.Command, args []string) {
			utils.ErrHandler(handlers.List(showAll, descending, listCount))
		},
	}
)

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "show all tasks.")
	listCmd.Flags().IntVarP(&listCount, "listcount", "c", 0, "Number of items to list (0 for all)")
	listCmd.Flags().BoolVarP(&descending, "descending", "d", false, "show records descending")
}
