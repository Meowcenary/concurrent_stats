package bootstrapregression

import (
	// "fmt"
	// "log"
	"math/rand"
	// "time"

	// "gonum.org/v1/gonum/stat"
)

type UnmappableError struct{
}

func (i *UnmappableError) Error() string {
	return "Unable to map indepdent variables to dependent variables for Point\nBe sure the slices passed are of the same length"
}

// struct to hold independent variable, dependent variable pair for
type Point struct {
	X float64
	Y float64
}

// Constructor
func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// if going with the original data format this will be something like
// {
//   field: [float1, float2, etc]
//
// func BootstrapRegression(data [string]float64) {
// 	// Initialize randomness
// 	seed := time.Now().UnixNano()
// 	fmt.Println("Seeding with: ", seed)
// 	rand.NewSource(seed)
//
// 	samples := createBootstrapSamples()
//
// 	// unassigned var in case of future refactors and for readability
// 	var weights []float64
// 	// Do not force the regression line to pass through the origin.
// 	origin := false
//
// 	// linear regression for each variable
// 	// ...
// }

// perform linear regression for a set of data
// func linearRegressionForSamples(samples []float64) {
// 	alpha, beta := stat.LinearRegression(xs, ys, nil, origin)
// 	r2 := stat.RSquared(xs, ys, nil, alpha, beta)
// 	// fmt.Printf("Estimated offset is: %.6f\n", alpha)
// 	// fmt.Printf("Estimated slope is:  %.6f\n", beta)
// 	// fmt.Printf("R^2: %.6f\n", r2)
// }

func CreatePoints(independentVarData []float64, dependentVarData []float64) ([]Point, error) {
	independentVarLen := len(independentVarData)
	dependentVarLen := len(dependentVarData)
	points := make([]Point, independentVarLen)

	if independentVarLen == dependentVarLen {
		for i := 0; i < independentVarLen; i++ {
			points[i] = *NewPoint(independentVarData[i], dependentVarData[i])
		}
	} else if independentVarLen != dependentVarLen {
		return nil, &UnmappableError{}
	}

	return points, nil
}

func CreateBootstrapSamples(r *rand.Rand, data []float64) [][]float64 {
	// the initial sample from which the bootstrap samples will to return are drawn
	initialSample := CreateBootstrapSample(r, data)

	// number of bootstrap samples to create
	samples := 1000
	bootstrapSamples := make([][]float64, samples)
	for i := 0; i < samples; i++ {
		// initialSample is now selected from instead of the original data
		bootstrapSamples[i] = CreateBootstrapSample(r, initialSample)
	}

	return bootstrapSamples
}

// the first argument is a pointer to rand.Rand so that the same random seed can be used all tests
// there is probably a better way to do this, but given time constraints this was the approach I went with
func CreateBootstrapSample(r *rand.Rand, data []float64) []float64 {
	sampleSize := len(data)
	bootstrapSample := make([]float64, sampleSize)

	for i := 0; i < sampleSize; i++ {
		bootstrapSample[i] = data[r.Intn(sampleSize)]
	}

	return bootstrapSample
}
