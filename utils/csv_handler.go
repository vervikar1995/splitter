package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

func handleError(context string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %v", context, err)
	}
	return nil
}

// Float64 custom type for handling float64 values
type Float64 float64

// UnmarshalCSV converts a CSV string to a Float64 type, replacing commas with dots
func (f *Float64) UnmarshalCSV(csvStr string) error {
	cleanedStr := strings.ReplaceAll(csvStr, ",", ".")
	value, err := strconv.ParseFloat(cleanedStr, 64)
	if err != nil {
		return handleError(fmt.Sprintf("error parsing number from string %s", csvStr), err)
	}
	*f = Float64(value)
	return nil
}

// DateTime is a custom type for handling date and time values
type DateTime struct {
	time.Time
}

// UnmarshalCSV converts a CSV date string to a DateTime type
func (date *DateTime) UnmarshalCSV(csvStr string) error {
	formats := []string{
		"02/01/2006", // DD/MM/YYYY
	}

	// Try parsing the date using the defined formats
	for _, format := range formats {
		parsedTime, err := time.Parse(format, csvStr)
		if err == nil {
			date.Time = parsedTime
			return nil
		}
	}
	return handleError(fmt.Sprintf("error parsing date %v", csvStr), nil)
}

// NoDateTime is a custom type for handling non-date time strings.
type NoDateTime struct {
	string
}

// UnmarshalCSV converts a CSV string to a NoDateTime type, formatting dates as necessary
func (date *NoDateTime) UnmarshalCSV(csvStr string) error {
	if strings.Contains(csvStr, "/") {
		// If the string contains a date format, parse and format it
		parsedTime, err := time.Parse("02/01/2006", csvStr)
		if err != nil {
			return handleError("error parsing date", err)
		}
		date.string = parsedTime.Format("2006-01-02") // Format to YYYY-MM-DD
	} else {
		date.string = csvStr // Preserve the original string format
	}
	return nil
}

// Table represents the structure of a row in the CSV file.
type Table struct {
	PersonName string     `csv:"person_name"`
	ID         int        `csv:"id"`
	Total      Float64    `csv:"Total"`
	Paid       Float64    `csv:"Paid"`
	Date       DateTime   `csv:"Date"`
	No         NoDateTime `csv:"No"`
}

// ReadCSVFile reads a CSV file and unmarshals its content into a slice of Table.
func ReadCSVFile(filePath string) ([]Table, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, handleError(fmt.Sprintf("error opening file %s", filePath), err)
	}
	defer file.Close()

	var tables []Table

	// Set the CSV reader with a custom delimiter
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';' // Use semicolon as the delimiter
		return r
	})

	// Unmarshal the file content into the tables slice
	if err := gocsv.Unmarshal(file, &tables); err != nil {
		return nil, handleError("error unmarshalling CSV", err)
	}

	return tables, nil
}
