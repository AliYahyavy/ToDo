package handlers

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"todo/utils"
)

func Search(id int, task string) {
	file, reader := utils.OpenFile("data.csv")
	flag := false
	defer file.Close()

	if id == 0 && task == "" {
		log.Fatal("Error: You must provide either the --id (-i) flag or the --task (-t) flag.")
	}

	for {
		record, readerErr := reader.Read()
		if readerErr == io.EOF {
			return
		}
		utils.ErrHandler(readerErr)
		if record[0] == strconv.Itoa(id) || record[1] == task {
			if !flag {
				fmt.Printf("ID   %-52s %-36s %-6s\n", "Task", "LastModify", "Done")
				flag = true
			}
			fmt.Printf("%-4s %-52s %-36s %-6s\n", record[0], record[1], record[2], record[3])
		}
	}
}
