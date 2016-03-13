package algorithms

import (
	"errors"
	"reflect"
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	testData := []struct {
		start          int
		end            int
		expectedError  error
		expectedResult int
	}{
		{0, 0, errors.New("error: start must be less than end"), 0},
		{9, 2, errors.New("error: start must be less than end"), 0},
		{1, 2, nil, 2},
		{1, 10, nil, 2520},
	}
	for _, data := range testData {
		result, _ := SmallestMultiple(data.start, data.end)
		if !reflect.DeepEqual(result, data.expectedResult) {
			t.Errorf("largestCommonFactors(%d, %d) failed. Expected %v. Got %v",
				data.start, data.end, data.expectedResult, result)
		}
	}
}

func TestLargestCommonFactors(t *testing.T) {
	testData := []struct {
		start          int
		end            int
		expectedError  error
		expectedResult *[]int
	}{
		{0, 0, errors.New("error: start must be less than end"), nil},
		{9, 2, errors.New("error: start must be less than end"), nil},
		{1, 2, nil, &[]int{1, 2}},
		{1, 10, nil, &[]int{1, 8, 9, 5, 7}},
	}
	for _, data := range testData {
		result, _ := largestCommonFactors(data.start, data.end)
		if !reflect.DeepEqual(result, data.expectedResult) {
			t.Errorf("largestCommonFactors(%d, %d) failed. Expected %v. Got %v",
				data.start, data.end, data.expectedResult, result)
		}
	}
}

func TestLeastCommonFactors(t *testing.T) {
	testData := []struct {
		start          int
		end            int
		expectedError  error
		expectedResult *[]int
	}{
		{0, 0, errors.New("error: start must be less than end"), nil},
		{9, 2, errors.New("error: start must be less than end"), nil},
		{1, 2, nil, &[]int{1, 2}},
		{1, 10, nil, &[]int{1, 2, 3, 5, 7}},
	}

	for _, data := range testData {
		result, err := leastCommonFactors(data.start, data.end)
		if !reflect.DeepEqual(result, data.expectedResult) {
			t.Errorf("leastCommonFactors(%d, %d) failed. Expected %v. Got %v",
				data.start, data.end, data.expectedResult, result)
		}
		if data.expectedError == nil && err != nil {
			t.Errorf("leastCommonFactors(%d, %d) failed with unexpected error: %v", data.start, data.end, err)
		}
		if data.expectedError != nil && !reflect.DeepEqual(data.expectedError, err) {
			t.Errorf("leastCommonFactors(%d, %d) failed. Expected error: %v. Got error: %v",
				data.start, data.end, data.expectedError, err)
		}
	}
}

func TestContainsFactorOf(t *testing.T) {
	testData := []struct {
		list           *[]int
		n              int
		expectedResult bool
	}{
		{&[]int{}, 1, false},
		{&[]int{0}, 1, false},
		{&[]int{0, 1, 2, 3}, 0, false},
		{&[]int{3, 4, 5}, 0, false},
		{&[]int{2, 4, 6, 8, 9}, 4, true},
		{&[]int{1, 3, 10, 12}, 40, true},
		{&[]int{8, 9, 7, 6, 5}, 2, false},
	}

	for _, data := range testData {
		if result := containsFactorOf(data.list, data.n); result != data.expectedResult {
			t.Errorf("containsFactorOf(%v, %d) failed. Expected %v. Got %v",
				data.list, data.n, data.expectedResult, result)

		}
	}
}
