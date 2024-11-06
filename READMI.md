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
```
git clone https://github.com/vervikar1995/splitter.git
cd splitter



2. Install dependencies:
```
  go get

3. Usage

Place your input CSV file named Test_task_1.csv in the root directory of the project.
Open your command line interface.
Navigate to the project directory.
Run the program using the following command:
    ```
        go run main.go

The output CSV files will be created in the output_files directory, with each file named `output_&lt;id&gt;`.csv.

4. Example Input
    ```
            person_name;id;Total;Paid;Date;No
            John Doe;1;4,5;2,5;21/08/2019;20180615
            Jane Smith;2;10,0;5,0;15/06/2018;15/06/2018

5. Example Output
            The program will create output files like:
    ```
            output_1.csv
            output_2.csv

    Each output file will contain formatted data according to the specified rules