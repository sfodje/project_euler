package algorithms_test

import (
	"testing"

	"github.com/sfodje/project_euler/algorithms"
)

func TestSumOfEvenFibonnaci(t *testing.T) {
	testData := make(map[int64]int64)
	testData[0] = 0
	testData[1] = 0
	testData[2] = 2
	testData[3] = 2
	testData[4] = 2
	testData[5] = 2
	testData[7] = 2
	testData[8] = 10
	testData[4000000] = 4613732

	for n, expectedValue := range testData {
		if value := algorithms.SumOfEvenFibonacci(n); value != expectedValue {
			t.Errorf("SumOfEvenFibonacci(%d) failed. Got %d, Expected %d", n, value, expectedValue)
		}
	}
}
