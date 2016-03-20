package algorithms

import (
	"reflect"
	"testing"
)

func TestGetMaxProduct(t *testing.T) {
	testData := []struct {
		n              int
		expectedResult int
	}{
		{0, 0},
		{1, 99},
		{2, 9603},
		{3, 811502},
		{4, 70600674},
	}

	for _, data := range testData {
		if result := GetMaxProduct(data.n); result != data.expectedResult {
			t.Errorf("GetMaxProduct(%d) failed. Expected result %d. Got result %d", data.n, data.expectedResult, result)
		}
	}
}

func TestMaxProduct(t *testing.T) {
	testData := []struct {
		array          *[]int
		n              int
		expectedResult int
	}{
		{&[]int{1, 2, 3, 4, 5, 6}, -1, 0},
		{&[]int{1, 2, 3, 4, 5, 6}, 0, 0},
		{&[]int{1, 2, 3, 4, 5, 6}, 1, 6},
		{&[]int{1, 2, 3, 4, 5, 6, 3, 1}, 2, 30},
		{&[]int{6, 5, 4, 3, 2, 1}, 3, 120},
	}

	for _, data := range testData {
		if result := maxProduct(data.array, data.n); result != data.expectedResult {
			t.Errorf("maxProduct(%v, %d) failed. Expected result %d. Got result %d",
				data.array, data.n, data.expectedResult, result)
		}
	}
}

func TestConvertGrid(t *testing.T) {
	testData := []struct {
		strgrid        string
		expectedResult *[][]int
	}{
		{"", nil},
		{"01", &[][]int{[]int{1}}},
		{"09 05 34 87 01 64", &[][]int{[]int{9, 5, 34, 87, 1, 64}}},
		{"09\n", &[][]int{[]int{9}}},
		{"\n02", &[][]int{[]int{2}}},
		{"09\n02", &[][]int{[]int{9}, []int{2}}},
		{"09 05 34\n04 87 23", &[][]int{[]int{9, 5, 34}, []int{4, 87, 23}}},
	}

	for _, data := range testData {
		if result := convertGrid(data.strgrid); !reflect.DeepEqual(result, data.expectedResult) {
			t.Errorf("convertGrid(%s) failed. Expected result %v. Got result %v", data.strgrid, data.expectedResult, result)
		}
	}
}
