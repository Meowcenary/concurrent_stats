package crossvalidationregression

import (
	"reflect"
	"testing"
)

func TestLeaveOutOneRegression(t *testing.T) {
	independentData := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	dependentData := []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}
	expectedAlpha := 10.0
	expectedBeta := 1.0

	alpha, beta := LeaveOutOneRegression(independentData, dependentData)
	if alpha != expectedAlpha {
		t.Errorf("Alpha does not match what was expected.\nExpected: %f\nGot: %f\n", expectedAlpha, alpha)
	} else if beta != expectedBeta {
		t.Errorf("Beta does not match waht was expected.\nExpected: %f\nGot: %f\n", expectedBeta, beta)
	}
}

func TestRemoveIndex(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	indicies := []int{0, 2, 4}
	expected := [][]float64{
		[]float64{2.0, 3.0, 4.0, 5.0},
		[]float64{1.0, 2.0, 4.0, 5.0},
		[]float64{1.0, 2.0, 3.0, 4.0},
	}

	for i := 0; i < len(indicies); i++ {
		result := RemoveIndex(data, indicies[i])
		expectedValue := expected[i]

		if !reflect.DeepEqual(result, expectedValue) {
			t.Errorf("Expected: %f\nGot: %f", expectedValue, result)
		}
	}
}
