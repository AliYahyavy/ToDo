package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Delete(id int) bool {

	if id == 0 {
		log.Fatal("Error: record ID (--id) is required for delete")
	}

	filePath := "data.csv"
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Error: file '%s' does not exist", filePath)
		}
		log.Fatalf("Error: failed to open file for reading: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	allRecords, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error: reading all CSV records failed: %v", err)
	}

	var recordsToKeep [][]string
	recordDeleted := false

	for _, record := range allRecords {

		recordIDStr := record[0]
		currentID, err := strconv.Atoi(recordIDStr)
		if err != nil {

			log.Printf("Warning: Skipping record with non-integer ID '%s' during delete check: %v\n", recordIDStr, err)
			recordsToKeep = append(recordsToKeep, record)
			continue
		}

		if currentID == id {
			recordDeleted = true
			continue
		}

		recordsToKeep = append(recordsToKeep, record)
	}

	if !recordDeleted {
		fmt.Println("Record not found with the specified ID.")
		return false
	}

	outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error: failed to open file for writing: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	err = writer.WriteAll(recordsToKeep)
	if err != nil {
		log.Fatalf("Error: writing updated CSV records failed: %v", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalf("Error: flushing CSV writer failed: %v", err)
	}

	return true
}
