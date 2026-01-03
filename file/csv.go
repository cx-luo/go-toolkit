// Package file provides file operation utilities
package file

import (
	"encoding/csv"
	"os"
)

// ReadCSV reads all records from a CSV file
func ReadCSV(csvFilePath string) ([][]string, error) {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// WriteCSV writes records to a CSV file
func WriteCSV(csvFilePath string, records [][]string) error {
	file, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(records)
}
