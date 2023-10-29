package crossvalidationregression

import (
	"reflect"
	"testing"
)

func TestRemoveIndex(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	indicies := []int{0, 2, 4}
	expected := [][]float64{
		[]float64{2.0, 3.0, 4.0, 5.0},
		[]float64{1.0, 2.0, 4.0, 5.0},
		[]float64{1.0, 2.0, 3.0, 4.0},
	}

	for i := 0; i < len(indicies); i++ {
		if !reflect.DeepEqual(RemoveIndex(data, indicies[i]), expected[i]) {
			t.Errorf("")
		}
	}
}
