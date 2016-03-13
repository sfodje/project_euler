package algorithms

import "testing"

func TestSumSquareDifference(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 4},
		{3, 22},
		{4, 70},
		{10, 2640},
		{100, 25164150},
	}

	for _, data := range testData {
		if result := SumSquareDifference(data.n); result != data.expectedResult {
			t.Errorf("SumSquareDifference(%d) failed. Expected %d. Got %d", data.n, data.expectedResult, result)
		}
	}
}

func TestSquareOfSum(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{2, 9},
		{3, 36},
		{4, 100},
	}

	for _, data := range testData {
		if result := squareOfSum(data.n); result != data.expectedResult {
			t.Errorf("squareOfSum(%d) failed. Expected %d. Got %d", data.n, data.expectedResult, result)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{-5, 0},
		{0, 0},
		{1, 1},
		{2, 5},
		{3, 14},
		{4, 30},
	}

	for _, data := range testData {
		if result := sumOfSquares(data.n); result != data.expectedResult {
			t.Errorf("sumOfSquares(%d) failed. Expected %d. Got %d", data.n, data.expectedResult, result)
		}
	}
}
