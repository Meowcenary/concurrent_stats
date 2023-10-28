package csvparser

import (
	"bufio"
	"encoding/csv"
	"os"
	"path"
)

type InvalidCsvPathError struct{
	Filepath string
}

func (i *InvalidCsvPathError) Error() string {
	return "Invalid CSV path: " + i.Filepath
}

// The proper way to handle this would be to use the csv library as in ReadCSV, but
// csv writer expects [][]string type which would mean refactoring things in a way
// that there is not time for. Best alternative was to write a formatted string
func WriteCSV(filepath string, data string) error {
	// Ensure path is a csv file
	if path.Ext(filepath) != ".csv" {
		return &InvalidCsvPathError{Filepath: filepath}
	}

	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(data)
	writer.Flush()

	return nil
}

func ReadCSV(filepath string) ([][]string, error) {
		// Ensure path is a csv file
		if path.Ext(filepath) != ".csv" {
			return nil, &InvalidCsvPathError{Filepath: filepath}
		}
		// Open file
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
    }
		// Defer keyword allows close call to be declared next to open call, but delays execution to end of function
		defer file.Close()

		// Read records from file
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
    if err != nil {
			return nil, err
    }

		return records, err
}

// create mapping of CSV headers to all values in column
// data is rows of CSV data read from a file with headers
func CsvDataByColumn(data [][]string) (map[string][]string, error) {
	headerIndex := make(map[int]string)
	dataByColumn := make(map[string][]string)

	// Pop header data off of data and create mapping to csv data
	headers, data := data[0], data[1:]
	for i, header := range headers {
		headerIndex[i] = header
	}

	for _, row := range data {
		for i, v := range row {
			header := headerIndex[i]
			dataByColumn[header] = append(dataByColumn[header], v)
		}
	}

	return dataByColumn, nil
}
