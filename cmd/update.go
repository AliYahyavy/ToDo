package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"todo/handlers"
	"todo/utils"
)

var (
	modifyID          int
	description, done string
	updateCmd         = &cobra.Command{
		Use:   "update",
		Short: "change columns value",
		Long:  "change task description or done column value",
		Run: func(cmd *cobra.Command, args []string) {
			affcted, err := handlers.Update(modifyID, description, done)
			utils.ErrHandler(err)
			if affcted {
				fmt.Printf("Task '%d' updtaed successfully.\n", modifyID)
			}
		},
	}
)

func init() {
	updateCmd.Flags().StringVarP(&description, "task", "t", "", "modify task description")
	updateCmd.Flags().StringVarP(&done, "done", "d", "", "modify task status")
	updateCmd.Flags().IntVarP(&modifyID, "id", "i", 0, "select task with id")
}
