package cmd

import (
	"github.com/spf13/cobra"
	"todo/handlers"
)

var (
	id   int
	task string

	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "Search specific task.",
		Long:  "Search record which can specified with id or task flag. it can return several records if use both flags at same time or several records have same task description",
		Run: func(cmd *cobra.Command, args []string) {
			handlers.Search(id, task)
		},
	}
)

func init() {
	searchCmd.Flags().IntVarP(&id, "id", "i", 0, "Search specific id")
	searchCmd.Flags().StringVarP(&task, "task", "t", "", "Search specific task")
	//searchCmd.MarkFlagsMutuallyExclusive("id", "task")
}
