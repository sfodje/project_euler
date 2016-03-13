package algorithms

import (
	"errors"
	"reflect"
	"testing"
)

func TestGreatestProductOfNDigits(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{0, 0},
		{1, 9},
		{4, 5832},
		{13, 23514624000},
	}
	for _, data := range testData {
		if result := GreatestProductOfNDigits(data.n); result != data.expectedResult {
			t.Errorf("GreatestProductOfNDigits(%d) failed. Expected %d. Got %d\n", data.n, data.expectedResult, result)
		}
	}
}

func TestProduct(t *testing.T) {
	testData := []struct {
		arr            *[]int
		expectedResult int
	}{
		{nil, 0},
		{&[]int{0}, 0},
		{&[]int{1}, 1},
		{&[]int{5}, 5},
		{&[]int{1, 2}, 2},
		{&[]int{3, 5}, 15},
		{&[]int{2, 3, 4, 5}, 120},
	}
	for _, data := range testData {
		if result := product(data.arr); result != data.expectedResult {
			t.Errorf("product(%v) failed. Expected result: %d. Got result: %d", data.arr, data.expectedResult, result)
		}
	}
}

func TestStringToIntArr(t *testing.T) {
	testData := []struct {
		str            string
		expectedResult *[]int
		expectedError  error
	}{
		{"1a", nil, errors.New("strconv.ParseInt: parsing \"a\": invalid syntax")},
		{"0", &[]int{0}, nil},
		{"12345", &[]int{1, 2, 3, 4, 5}, nil},
	}

	for _, data := range testData {
		result, err := stringToIntArray(data.str)
		if data.expectedError == nil && err != nil {
			t.Errorf("stringToIntArr(%s) failed unexpectedly. Error %v", data.str, err)
		}
		if data.expectedError != nil && err != nil && data.expectedError.Error() != err.Error() {
			t.Errorf("stringToIntArr(%s) failed. Expected error: %v. Got error: %v", data.str, data.expectedError, err)
		}
		if !reflect.DeepEqual(result, data.expectedResult) {
			t.Errorf("stringToIntArr(%s) failed. Expected result: %v. Got result: %v", data.str, data.expectedResult, result)
		}
	}
}
