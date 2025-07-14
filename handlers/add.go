package handlers

import (
	"encoding/csv"
	"os"
	"strconv"
	time2 "time"
	"todo/utils"
)

//type Task struct {
//	ID          string
//	Description string
//	CreatedAt   string
//	IsComplete  string
//}

func /*(t *Task)*/ Add(Description string, isDone bool) bool {
	time := time2.Now()
	iso8601Format := time.Format("2006-01-02T15:04:05-07:00")
	isDoneH := "No"
	if strconv.FormatBool(isDone) == "true" {
		isDoneH = "Yes"
	}
	task := []string{utils.RowCount(), Description, iso8601Format, isDoneH}
	file, openFileErr := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	utils.ErrHandler(openFileErr)
	defer file.Close()

	writer := csv.NewWriter(file)
	writeCsvErr := writer.Write(task)
	utils.ErrHandler(writeCsvErr)

	writer.Flush()

	return true
}
