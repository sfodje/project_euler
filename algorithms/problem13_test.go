package algorithms

import "testing"

func TestFirstNthDigitsOfSum(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int64
	}{
		{-1, 0},
		{0, 0},
		{1, 5},
		{2, 55},
		{3, 553},
		{4, 5537},
		{5, 55373},
		{6, 553737},
		{7, 5537376},
		{8, 55373762},
		{9, 553737623},
		{10, 5537376230},
	}

	for _, data := range testData {
		if result := FirstNthDigitsOfSum(data.n); result != data.expectedResult {
			t.Errorf("firstNthDigitsOfSum(%d) failed. Expected result %d. Got result %d.", data.n, data.expectedResult, result)
		}
	}
}
