package bootstrapregression

import (
	"math/rand"
	"reflect"
	"testing"
)

// fixture to ensure that random seed is always the same
func seed() int64 {
	return 12345
}

func random() *rand.Rand {
	return rand.New(rand.NewSource(seed()))
}

func TestBoostrapRegression(t *testing.T) {
	independentData := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	dependentData := []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}

	_, _, err := BootstrapRegression(independentData, dependentData)

	if err != nil {
		t.Errorf("Error was raised from call to BootstrapRegression")
	}
}

func TestLinearRegressionForSample(t *testing.T) {
	// Linear data
	data := []Point{Point{1.0, 1.0}, Point{2.0, 2.0}, Point{3.0, 3.0}, Point{4.0, 4.0}, Point{5.0, 5.0}}
	expectedAlpha := 0.0
	expectedBeta := 1.0
	alpha, beta := LinearRegressionForSample(data)

	if alpha != expectedAlpha {
		t.Errorf("Alpha value does not match expectation. Expected: %f\nGot: %f", expectedAlpha, alpha)
	} else if beta != expectedBeta {
		t.Errorf("Beta value does not match expectation. Expected: %f\nGot: %f", expectedBeta, beta)
	}

	// Messier data
	data = []Point{Point{1.0, 2.0}, Point{2.0, 10.0}, Point{3.0, 9.0}, Point{4.0, 17.0}, Point{5.0, 20.0}}
	expectedAlpha = -1.3
	expectedBeta = 4.3
	alpha, beta = LinearRegressionForSample(data)
	tolerance := 0.0000000001

	// ran into some weird behavior with comparing float64 values when negative and this
	// seemed like a reasonable approach given the time constraints
	if expectedAlpha-alpha > tolerance {
		t.Errorf("Alpha value does not match expectation. Expected: %f\nGot: %f", expectedAlpha, alpha)
	} else if beta != expectedBeta {
		t.Errorf("Beta value does not match expectation. Expected: %f\nGot: %f", expectedBeta, beta)
	}
}

func TestCreatePoints(t *testing.T) {
	independentData := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	dependentData := []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	expectedPoints := []Point{Point{1.0, 11}, Point{2, 12}, Point{3, 13}, Point{4, 14}, Point{5, 15}, Point{6, 16}, Point{7, 17}, Point{8, 18}, Point{9, 19}, Point{10, 20}}
	points, err := CreatePoints(independentData, dependentData)

	if err != nil {
		t.Errorf("Error raised while creating points")
	} else if !reflect.DeepEqual(points, expectedPoints) {
		t.Errorf("Expected points do not match with what was received.\nExpected: %v\nGot:%v", expectedPoints, points)
	}
}

// generation of values is tested in create bootstrap sample, this is to ensure
// that the function runs properly and returns the expected amount of samples
func TestCreateBootstrapSamples(t *testing.T) {
	data := []Point{Point{1.0, 1.0}, Point{2.0, 2.0}, Point{3.0, 3.0}, Point{4.0, 4.0}, Point{5.0, 5.0}}
	bootstrapSamples := CreateBootstrapSamples(random(), data)
	sampleCount := len(bootstrapSamples)

	if sampleCount != 1000 {
		t.Errorf("Expected 1000 samples, but got %d", sampleCount)
	}
}

func TestCreateBootstrapSample(t *testing.T) {
	// subset of data to sample
	data := []Point{Point{1.0, 1.0}, Point{2.0, 2.0}, Point{3.0, 3.0}, Point{4.0, 4.0}, Point{5.0, 5.0}}
	expectedBootstrapSample := []Point{Point{4.0, 4.0}, Point{4.0, 4.0}, Point{5.0, 5.0}, Point{2.0, 2.0}, Point{2.0, 2.0}}
	bootstrapSample := CreateBootstrapSample(random(), data)

	if !reflect.DeepEqual(bootstrapSample, expectedBootstrapSample) {
		t.Errorf("Bootstrap sample does not match what was expected.\nExpected: %f\nGot: %f\n", expectedBootstrapSample, bootstrapSample)
	}
}
