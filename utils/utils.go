package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// SplitCSVFile takes a slice of Table structs and splits them into separate CSV files based on unique IDs
func SplitCSVFile(tables []Table) error {
	// Create a map to hold unique records indexed by ID
	uniqueRecords := make(map[int][][]string)

	// Iterate over each table and prepare records for CSV
	for _, table := range tables {
		record := []string{
			table.PersonName,
			strconv.Itoa(table.ID),
			strconv.FormatFloat(float64(table.Total), 'f', 2, 64),
			strconv.FormatFloat(float64(table.Paid), 'f', 2, 64),
			table.Date.Format("02-Jan-06"), // Format the date
			table.No.string,
		}

		// Append the record to the corresponding ID in the map
		uniqueRecords[table.ID] = append(uniqueRecords[table.ID], record)
	}

	outputDir := "output_files"
	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory %s: %v", outputDir, err)
	}

	// Write each unique ID's records to a separate CSV file
	for id, records := range uniqueRecords {
		outputFilePath := filepath.Join(outputDir, fmt.Sprintf("output_%d.csv", id))
		if err := writeCSVFile(outputFilePath, records); err != nil {
			return fmt.Errorf("error writing to file %s: %v", outputFilePath, err)
		}
	}

	return nil
}

// writeCSVFile writes the given records to a CSV file specified by the filePath.
func writeCSVFile(filePath string, records [][]string) error {
	// Create a new CSV file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	writer := csv.NewWriter(file) // Create a new CSV writer
	defer writer.Flush()          // Flush any buffered data to the underlying writer

	// Write the header row to the CSV file
	if err := writer.Write([]string{"person_name", "id", "Total", "Paid", "Date", "No"}); err != nil {
		return fmt.Errorf("error writing header: %v", err)
	}

	// Write each record to the CSV file
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record: %v", err)
		}
	}

	return nil
}
