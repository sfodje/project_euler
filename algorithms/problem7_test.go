package algorithms

import "testing"

func TestNthPrimeNumber(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 11},
		{6, 13},
		{7, 17},
		{10001, 104743},
	}

	for _, data := range testData {
		if result := NthPrimeNumber(data.n); result != data.expectedResult {
			t.Errorf("NthPrimeNumber(%d) failed. Expected %d. Got %d", data.n, data.expectedResult, result)
		}
	}
}
