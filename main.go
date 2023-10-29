package main

import (
	"fmt"
	"log"

	"github.com/Meowcenary/concurrent_stats/csvparser"
	// "github.com/Meowcenary/concurrent_stats/boostrapregression"
)

func main() {
	// Read data from CSV
	file := "boston.csv"
	records, err := csvparser.ReadCSV(file)
	if err != nil {
		log.Fatal(err)
	}

	// Parse data from CSV
	data, err := csvparser.CsvDataByColumn(records)
	if err != nil {
		log.Fatal(err)
	}
}
