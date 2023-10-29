package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Meowcenary/concurrent_stats/csvparser"
	"github.com/Meowcenary/concurrent_stats/bootstrapregression"
	"github.com/Meowcenary/concurrent_stats/crossvalidationregression"
)

// main() can only be defined once, but two functions need to be executable separately of each other
// this approach solves the issue, but ideally this would be separated out into a sub package that was
// itself exectuable. Given time contraints this wasn't done. An attempt was made to define two files
// concurrent_regression.go and serial_regression.go each with a main() function and this worked in that
// each program could be compiled and ran, but tests would no longer work due to the double definition of main()
func main() {
	// CLI args
	args := os.Args
	argumentError := "This program expects a single command line argument of either \"serial\" or \"concurrent\" without quotes"

	if len(args) != 2 {
		fmt.Println(argumentError)
	} else if args[1] == "serial" {
		SerialRegression()
	} else if args[1] == "concurrent" {
		ConcurrentRegression()
	} else {
		fmt.Println(argumentError)
	}
}

func SerialRegression() {
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

	fmt.Printf("Bootstrap Regression\n---\n")
	for _, explanatoryVar := range explanatoryVars {
		alpha, beta, err := bootstrapregression.BootstrapRegression(data[explanatoryVar], data[responseVar])

		if err != nil {
			break
		}

		fmt.Printf("Coefficients for explanatory variable %s\nAlpha: %f, Beta: %f\n", explanatoryVar, alpha, beta)
	}

	fmt.Printf("\nLeave Out One Cross Validation Regression\n---\n")
	for _, explanatoryVar := range explanatoryVars {
		alpha, beta := crossvalidationregression.LeaveOutOneRegression(data[explanatoryVar], data[responseVar])

		fmt.Printf("Coefficients for explanatory variable %s\nAlpha: %f, Beta: %f\n", explanatoryVar, alpha, beta)
	}
}

func ConcurrentRegression() {
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
