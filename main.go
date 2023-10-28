package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat"

	"github.com/Meowcenary/concurrent_stats/csvparser"
)

func main() {
	// Initialize randomness
	seed := time.Now().UnixNano()
	fmt.Println("Seeding with: ", seed)
	rand.NewSource(seed)

	// Read data from CSV
	file := "datafilepath"
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

func bootstrapRegression() {
	// Create samples...
	bootstrapSamples := createBootstrapSamples(data)

	// naieve implementation to just get this moving
	for _, sample := range bootstrapSamples {
		for i, v := range sample {
		}

		// run linear regression on the data and average across the sets
		stat.LinearRegression(, , nil, false)
	}
}

// Create the bootstrap samples to be used for analysis
// returns a two dimensional array consisting of the bootstrap sets
// visualization of process: https://en.wikipedia.org/wiki/Bootstrapping_(statistics)#/media/File:Illustration_bootstrap.svg
// TODO: potential optimization would be to instantiate the bootstrapSamples slice to 1000 (or n) records immediately so that it
// would not need to grow the slice dynamically
func createBootstrapSamples(data map[string]float64) [][]float64 {
	sampleSize := len(data)
	// the initial sample from which the bootstrap samples will to return are drawn
	initialSample := createBootstrapSample(data)

	var bootstrapSamples [][]float64

	// 1000 is the number of bootstrap samples to create
	for i := 0; i < 1000; i++ {
		// initialSample is now selected from instead of the original data
		append(bootstrapSamples, createBootstrapSample(initialSample))
	}

	return bootstrapSamples
}

// pick a random number that is an index to the data array and add the value there to the bootstrap sample
// TODO: pass the data by reference to avoid copying
func createBootstrapSample(data map[string]float64) []float64 {
	var bootstrapSample []float64

	for i, v := range sample {
		append(bootstrapSample, data[rand.Intn(sampleSize)])
	}

	return bootstrapSample
}
