package main

import (
	"fmt"
	"log"

	"github.com/Meowcenary/concurrent_stats/csvparser"
	"github.com/Meowcenary/concurrent_stats/bootstrapregression"
	"github.com/Meowcenary/concurrent_stats/crossvalidationregression"
)

func main() {
	// channels to hold values from routines
	bootstrapChannel := make(chan float64)
	crossvalidationChannel := make(chan float64)

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

	go RunBoostrapRegression(data, explanatoryVars, responseVar, &bootstrapChannel)
	go RunCrossValidationRegression(data, explanatoryVars, responseVar, &crossvalidationChannel)

	fmt.Printf("Bootstrap Regression\n---\n")
	for i := 0; i < len(explanatoryVars); i++ {
		alpha, beta := <-bootstrapChannel, <-bootstrapChannel
		fmt.Printf("Coefficients for explanatory variable %s\nAlpha: %f, Beta: %f\n", explanatoryVars[i], alpha, beta)
	}

	fmt.Printf("\nLeave Out One Cross Validation Regression\n---\n")
	for i := 0; i < len(explanatoryVars); i++ {
		alpha, beta := <-crossvalidationChannel, <-crossvalidationChannel
		fmt.Printf("Coefficients for explanatory variable %s\nAlpha: %f, Beta: %f\n", explanatoryVars[i], alpha, beta)
	}
}

func RunBoostrapRegression(data map[string][]float64, explanatoryVars []string, responseVar string, channel *chan float64) {
	for _, explanatoryVar := range explanatoryVars {
		alpha, beta, err := bootstrapregression.BootstrapRegression(data[explanatoryVar], data[responseVar])

		if err != nil {
			break
		}

		*channel <- alpha
		*channel <- beta
	}
}

func RunCrossValidationRegression(data map[string][]float64, explanatoryVars []string, responseVar string, channel *chan float64) {
	for _, explanatoryVar := range explanatoryVars {
		alpha, beta := crossvalidationregression.LeaveOutOneRegression(data[explanatoryVar], data[responseVar])

		*channel <- alpha
		*channel <- beta
	}
}
