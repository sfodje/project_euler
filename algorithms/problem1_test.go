package algorithms_test

import (
	"testing"

	"github.com/sfodje/project_euler/algorithms"
)

func TestSumOfMultiplesOf3And5(t *testing.T) {
	testData := make(map[int]int)
	testData[0] = 0
	testData[1] = 0
	testData[2] = 0
	testData[3] = 0
	testData[4] = 3
	testData[5] = 3
	testData[6] = 8
	testData[10] = 23
	testData[1000] = 233168

	for n, expectedValue := range testData {
		if value := algorithms.SumOfMultiplesOf3And5(n); value != expectedValue {
			t.Errorf("SumOfMultiplesOf3And5(%d) failed (Expected %d, Got %d)", n, expectedValue, value)
		}
		if value := algorithms.SumOfMultiplesOf3And5Iterative(n); value != expectedValue {
			t.Errorf("SumOfMultiplesOf3And5(%d) failed (Expected %d, Got %d)", n, expectedValue, value)
		}
	}

}
