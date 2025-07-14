package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	time2 "time"
)

func Update(id int, task, isDone string) (bool, error) {
	if id == 0 {
		return false, fmt.Errorf("error: record ID (--id) is required for update")
	}
	if task == "" && isDone == "" {
		return false, fmt.Errorf("error: at least one of --task (-t) or --done (-d) must be provided to update")
	}
	if isDone != "" && isDone != "Yes" && isDone != "No" {
		return false, fmt.Errorf("error: --done (-d) flag must be either 'Yes' or 'No'")
	}

	filePath := "data.csv"
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("file '%s' does not exist", filePath)
		}
		return false, fmt.Errorf("failed to open file for reading: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	allRecords, err := reader.ReadAll()
	if err != nil {
		return false, fmt.Errorf("error reading all CSV records: %w", err)
	}

	recordFoundAndUpdated := false
	for i, record := range allRecords {
		recordIDStr := record[0]

		currentID, err := strconv.Atoi(recordIDStr)
		if err != nil {
			log.Printf("Warning: Skipping record with non-integer ID '%s': %v\n", recordIDStr, err)
			continue
		}

		if currentID == id {

			if task != "" {
				allRecords[i][1] = task
			}
			if isDone != "" {
				allRecords[i][3] = isDone
			}
			time := time2.Now()
			iso8601Format := time.Format("2006-01-02T15:04:05-07:00")
			allRecords[i][2] = iso8601Format
			recordFoundAndUpdated = true
			break
		}
	}

	if !recordFoundAndUpdated {
		fmt.Println("Record not find.")
		return false, nil
	}

	outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return false, fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	err = writer.WriteAll(allRecords)
	if err != nil {
		return false, fmt.Errorf("error writing updated CSV records: %w", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return false, fmt.Errorf("error flushing CSV writer: %w", err)
	}

	return true, nil
}
