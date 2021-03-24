package mymath

import "testing"

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		result float64
	}

	tests := []test{
		test{[]int{1, 4, 6, 8, 100}, 6},
		test{[]int{0, 8, 10, 1000}, 9},
	}

	for _, res := range tests {
		x := CenteredAvg(res.data)
		if x != res.result {
			t.Error("Expected: ", res.result, " got: ", x)
		}
	}
}
