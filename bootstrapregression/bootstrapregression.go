package bootstrapregression

import (
	// "fmt"
	// "log"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat"
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

func BootstrapRegression(independentData []float64, dependentData []float64) (alpha float64, beta float64, err error) {
	// Initialize randomness
	seed := time.Now().UnixNano()
	data, err := CreatePoints(independentData, dependentData)

	if err != nil {
		return 0.0, 0.0, err
	}

	samples := CreateBootstrapSamples(rand.New(rand.NewSource(seed)), data)
	totalAlpha := 0.0
	totalBeta := 0.0

	for _, sample := range samples {
		alpha, beta := LinearRegressionForSample(sample)
		totalAlpha += alpha
		totalBeta += beta
	}

	samplesSize := float64(len(samples))
	return totalAlpha/samplesSize, totalBeta/samplesSize, nil
}

// perform linear regression for sample of data returning
// alpha - the offset and beta - the slope coefficient
func LinearRegressionForSample(data []Point) (alpha, beta float64) {
	dataLength := len(data)
	independentVarData := make([]float64, dataLength)
	dependentVarData := make([]float64, dataLength)

	for i := 0; i < dataLength; i++ {
		independentVarData[i] = data[i].X
		dependentVarData[i] = data[i].Y
	}

	return stat.LinearRegression(independentVarData, dependentVarData, nil, false)
}

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

func CreateBootstrapSamples(r *rand.Rand, data []Point) [][]Point {
	// the initial sample from which the bootstrap samples will to return are drawn
	initialSample := CreateBootstrapSample(r, data)

	// number of bootstrap samples to create
	samples := 1000
	bootstrapSamples := make([][]Point, samples)
	for i := 0; i < samples; i++ {
		// initialSample is now selected from instead of the original data
		bootstrapSamples[i] = CreateBootstrapSample(r, initialSample)
	}

	return bootstrapSamples
}

// the first argument is a pointer to rand.Rand so that the same random seed can be used all tests
// there is probably a better way to do this, but given time constraints this was the approach I went with
func CreateBootstrapSample(r *rand.Rand, data []Point) []Point {
	sampleSize := len(data)
	bootstrapSample := make([]Point, sampleSize)

	for i := 0; i < sampleSize; i++ {
		bootstrapSample[i] = data[r.Intn(sampleSize)]
	}

	return bootstrapSample
}
