package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"todo/handlers"
	"todo/utils"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete desire record",
		Long:  "delete desire record with id",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			utils.ErrHandler(err)
			if affected := handlers.Delete(id); affected {
				fmt.Println("Record deleted successfully.")
			}
		},
	}
)
