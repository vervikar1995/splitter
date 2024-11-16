package main

import (
	"fmt"
	"log"
	"testTask/pkg"
)

func main() {
	// Define the path to the input CSV file
	filePath := "Test_task_1.csv"

	// Read the CSV file and unmarshal it into a slice of Table structs
	tables, err := pkg.ReadCSVFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err) // Log fatal error if reading fails
	}

	// Split the read tables into separate CSV files based on unique IDs
	if err := pkg.SplitCSVFile(tables); err != nil {
		log.Fatalf("Error splitting CSV file: %v", err) // Log fatal error if splitting fails
	}

	fmt.Println("Files are created. You can find them in the 'output_files' directory!")
}
