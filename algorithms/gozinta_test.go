package algorithms_test

import (
	"testing"

	"github.com/sfodje/project_euler/algorithms"
)

func TestSumOfGozintaChains(t *testing.T) {
	testData := make(map[int64]int64)
	testData[0] = 0
	testData[1] = 1
	testData[2] = 1
	testData[3] = 1
	testData[4] = 2
	testData[5] = 1
	testData[6] = 3
	testData[7] = 1
	testData[8] = 4
	testData[12] = 8
	testData[48] = 48
	testData[120] = 132
	for n, expectedValue := range testData {
		value := algorithms.NumberOfGozintaChains(n)
		if expectedValue != value {
			t.Errorf("NumberOfGozintaChains(%d) failed (expected %d, Got %d)", n, expectedValue, value)
		}
	}
}

func TestGozintaProblem(t *testing.T) {
	testData := make(map[int64]int64)
	testData[0] = 0
	testData[1] = 1
	testData[2] = 1
	testData[10] = 1
	testData[30] = 1
	testData[48] = 48 + 1
	testData[55] = 48 + 1
	testData[10000] = 3825
	for n, expectedValue := range testData {
		value := algorithms.GozintaProblem(n)
		if expectedValue != value {
			t.Errorf("GozintaProblem(%d) failed (expected %d, Got %d)", n, expectedValue, value)
		}
	}
}
