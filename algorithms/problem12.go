package algorithms

import "math"

func TriangleNumberWithMoreThanNDivisors(n int) int {
	if n < 0 {
		return 0
	}
	var i int
	factors := []int{}
	for j := 1; len(factors) <= n; j++ {
		i += j
		factors = *getFactors(i)
	}
	return i
}

func getFactors(n int) *[]int {
	if n == 0 || n == 1 {
		return &[]int{n}
	}
	factors := []int{1, n}
	for i := 2; i < int(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 {
			if i*i != n {
				factors = append(factors, i, n/i)
			} else {
				factors = append(factors, i)
			}
		}
	}
	return &factors
}
