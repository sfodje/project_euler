package algorithms

import (
	"errors"
	"reflect"
	"testing"
)

func TestProductOfSpecialPythTriplet(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult float64
	}{
		{1, 0},
		{1000, 31875000},
	}

	for _, data := range testData {
		if result := ProductOfSpecialPythTriplet(data.n); result != data.expectedResult {
			t.Errorf("ProductOfSpecialPythTriplet(%d) failed. Expected result: %v. Got result: %v", data.n, data.expectedResult, result)
		}
	}
}

func TestValueOfA(t *testing.T) {
	testData := []struct {
		b              int
		n              int
		expectedResult float64
		expectedError  error
	}{
		{1, 1, 0, errors.New("Error: b cannot be equal to n")},
		{7, 7, 0, errors.New("Error: b cannot be equal to n")},
		{1, 2, 0, nil},
		{3, 7, float64(0.875), nil},
		{7, 5, float64(11.25), nil},
	}

	for _, data := range testData {
		result, err := valueOfA(data.b, data.n)
		if data.expectedError == nil && err != nil {
			t.Errorf("valueOfA(%d, %d) failed unexpectedly. Error %v", data.b, data.n, err)
		}
		if data.expectedError != nil && err != nil && data.expectedError.Error() != err.Error() {
			t.Errorf("valueOfA(%d, %d) failed. Expected error: %v. Got error: %v", data.b, data.n, data.expectedError, err)
		}
		if !reflect.DeepEqual(result, data.expectedResult) {
			t.Errorf("valueOfA(%d, %d) failed. Expected result: %v. Got result: %v", data.b, data.n, data.expectedResult, result)
		}
	}
}

func TestIsAWholeNumber(t *testing.T) {
	testData := []struct {
		v              interface{}
		expectedResult bool
	}{
		{1, true},
		{1.1, false},
		{"1", false},
		{float64(1.23), false},
		{float32(9), true},
		{float64(7), true},
		{float64(9.3), false},
	}

	for _, data := range testData {
		if result := isWholeNumber(data.v); result != data.expectedResult {
			t.Errorf("isWholeNumber(%v) failed. Expected %v. Got %v", data.v, data.expectedResult, result)
		}
	}
}
