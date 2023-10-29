package main

import (
	"fmt"
	"log"

	"github.com/Meowcenary/concurrent_stats/csvparser"
	"github.com/Meowcenary/concurrent_stats/bootstrapregression"
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

	// Can be single column or multiple
	explanatoryVars := []string{"dis", "zn"}
	responseVar := "mv"

	for _, explanatoryVar := range explanatoryVars {
		alpha, beta, err := bootstrapregression.BootstrapRegression(data[explanatoryVar], data[responseVar])

		if err != nil {
			break
		}

		fmt.Printf("Coefficients for explanatory variable %s\nAlpha: %f, Beta: %f\n", explanatoryVar, alpha, beta)
	}
}
