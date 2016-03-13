package algorithms

import (
	"errors"
	"math"
)

func SmallestMultiple(start, end int) (int, error) {
	smallestMultiple := 1
	factors, err := largestCommonFactors(start, end)
	if err != nil {
		return 0, err
	}
	for _, f := range *factors {
		smallestMultiple *= f
	}
	return smallestMultiple, nil
}

func logN(n float64, base int) float32 {
	return float32(math.Log10(n) / math.Log10(float64(base)))
}

func largestCommonFactors(start, end int) (*[]int, error) {
	commonFactors, err := leastCommonFactors(start, end)
	if err != nil {
		return nil, err
	}
	for i, f := range *commonFactors {
		for j := end; j >= f*f; j-- {
			if m := logN(float64(j), f); m == float32(int32(m)) {
				(*commonFactors)[i] = j
				break
			}
		}
	}
	return commonFactors, nil
}

func leastCommonFactors(start, end int) (*[]int, error) {
	if !(start < end) {
		return nil, errors.New("error: start must be less than end")
	}
	var factors []int
	for i := start; i <= end; i++ {
		if !containsFactorOf(&factors, i) {
			factors = append(factors, i)
		}
	}
	return &factors, nil
}

func containsFactorOf(list *[]int, n int) bool {
	if list == nil || len(*list) == 0 || n == 0 {
		return false
	}
	for _, l := range *list {
		if l == 0 || l == 1 {
			continue
		}
		if n%l == 0 {
			return true
		}
	}
	return false
}
