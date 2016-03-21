package algorithms

import (
	"reflect"
	"testing"
)

func TestSumOfPrimes(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 2},
		{4, 5},
		{5, 5},
		{6, 10},
		{7, 10},
		{8, 17},
		{9, 17},
		{10, 17},
	}

	for _, data := range testData {
		if result := SumOfPrimes(data.n); result != data.expectedResult {
			t.Errorf("SumOfPrimes(%d) failed. Expected result: %d. Got result: %d", data.n, data.expectedResult, result)
		}
	}
}

func TestGetPrimes(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult *[]int
	}{
		{1, &[]int{}},
		{2, &[]int{2}},
		{3, &[]int{2, 3}},
		{4, &[]int{2, 3}},
		{5, &[]int{2, 3, 5}},
		{6, &[]int{2, 3, 5}},
		{7, &[]int{2, 3, 5, 7}},
		{100, &[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, data := range testData {
		if result := getPrimes(data.n); !reflect.DeepEqual(result, data.expectedResult) {
			t.Errorf("GetPrimes(%d) failed. Expected result: %v. Got result: %v", data.n, data.expectedResult, result)
		}
	}
}
