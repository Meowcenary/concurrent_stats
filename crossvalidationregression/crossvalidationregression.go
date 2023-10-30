package crossvalidationregression

import (
	"gonum.org/v1/gonum/stat"
)

// Cross validation that leaves out one record
// A good refactor would be to accept an integer argument p - the number of values to leave out
// Calculate regression for data with one value removed each time and take the average
func LeaveOutOneRegression(independentVarData []float64, dependentVarData []float64) (alpha float64, beta float64) {
	dataSize := len(independentVarData)
	totalAlpha := 0.0
	totalBeta := 0.0

	for i := 0; i < dataSize; i++ {
		// Remove value from data at index i
		independentDataForIteration := RemoveIndex(independentVarData, i)
		dependentDataForIteration := RemoveIndex(dependentVarData, i)

		// Calculate coefficients
		alpha, beta := stat.LinearRegression(independentDataForIteration, dependentDataForIteration, nil, false)

		// Add to total
		totalAlpha += alpha
		totalBeta += beta
	}

	return totalAlpha/float64(dataSize), totalBeta/float64(dataSize)
}

func RemoveIndex(s []float64, index int) []float64 {
    ret := make([]float64, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}
