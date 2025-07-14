package handlers

import (
	"fmt"
	"todo/utils"
)

func List(showAll, descending bool, listCount int) error {
	var displayedCount int

	file, reader := utils.OpenFile("data.csv")
	defer file.Close()
	allRecords, readerErr := reader.ReadAll()
	utils.ErrHandler(readerErr)

	fmt.Printf("ID   %-52s %-36s %-6s\n", "Task", "LastModify", "Done")

	if descending {
		descendingReadFile(allRecords, displayedCount, showAll, listCount)
		return nil
	}
	ascendingReadFile(allRecords, displayedCount, showAll, listCount)

	return nil
}

func ascendingReadFile(allRecords [][]string, displayedCount int, showAll bool, listCount int) {
	for _, row := range allRecords {
		if showAll || row[3] != "Yes" {
			if listCount > 0 && displayedCount >= listCount {
				break
			}
			fmt.Printf("%-4s %-52s %-36s %-6s\n", row[0], row[1], row[2], row[3])
			displayedCount++
		}
	}
}

func descendingReadFile(allRecords [][]string, displayedCount int, showAll bool, listCount int) {
	for i := len(allRecords) - 1; i >= 0; i-- {
		row := allRecords[i]

		if showAll || row[3] != "Yes" {
			if listCount > 0 && displayedCount >= listCount {
				break
			}
			fmt.Printf("%-4s %-52s %-36s %-6s\n", row[0], row[1], row[2], row[3])
			displayedCount++
		}
	}
}
