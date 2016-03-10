package algorithms

import (
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
	testData := make(map[int64]int64)
	testData[0] = 0
	testData[1] = 0
	testData[2] = 2
	testData[3] = 3
	testData[4] = 2
	testData[5] = 5
	testData[6] = 3
	testData[7] = 7
	testData[600851475143] = 6857
	for arg, expectedValue := range testData {
		if value := LargestPrimeFactor(arg); value != expectedValue {
			t.Errorf("LargestPrimeFactor(%d) failed. Expected %d, got %d", arg, expectedValue, value)
		}
	}
}

func TestIsPrime(t *testing.T) {
	testData := []struct {
		n              int64
		expectedResult bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{11, true},
		{13, true},
		{19, true},
		{6857, true},
	}

	for _, data := range testData {
		if result := isPrime(data.n); result != data.expectedResult {
			t.Errorf("isPrime(%d) failed. Expected %v. Got %v", data.n, data.expectedResult, result)
		}
	}
}
