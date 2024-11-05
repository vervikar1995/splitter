# CSV Splitter

A simple Go program that reads a CSV file and splits it into multiple CSV files based on unique IDs. This program processes the input data to ensure proper formatting and handles specific requirements for data representation.

## Features

- **Reads CSV Files**: Supports CSV files formatted with semicolons (`;`).
- **Unique ID Handling**: Splits records into separate files for each unique ID.
- **Custom Data Formatting**:
- Replaces commas with dots in numeric values.
- Converts date formats to `dd-MMM-yy`.
- Values in the "No" column are handled as follows:
- Values in the format `20180615` remain unchanged.
- Values in the format `dd/mm/yyyy` are reformatted to `yyyy-mm-dd` (e.g., `15/06/2018` becomes `2018-06-15`).
- The program does not alter date formats in the "No" column if they appear as `yyyy.mm.dd` (e.g., `2018.06.14` remains unchanged).
- **Error Handling**: Robust error handling for file operations and data processing.

## Requirements

- Go programming language (version 1.16 or higher)
- External library: `github.com/gocarina/gocsv` for CSV handling

## Installation

1. Clone the repository:
    git clone <repository-url>
    cd <repository-directory>


2. Install dependencies:
    go get
