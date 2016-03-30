package algorithms

import "testing"

func TestTriangleNumberWithMoreThanNDivisors(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{-1, 0},
		{0, 1},
		{1, 3},
		{2, 6},
		{4, 28},
		{500, 76576500},
	}

	for _, data := range testData {
		if result := TriangleNumberWithMoreThanNDivisors(data.n); result != data.expectedResult {
			t.Errorf("TriangleNumberWithMoreThanNDivisors(%d) failed. Expected result: %d. Got result: %d.",
				data.n, data.expectedResult, result)
		}
	}
}
