package algorithms

func SumSquareDifference(n int) int {
	return squareOfSum(n) - sumOfSquares(n)
}

func squareOfSum(n int) int {
	s := sum(n)
	return s * s
}

func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}

func sumOfSquares(n int) int {
	if n < 0 {
		return 0
	}
	if n < 1 {
		return n * n
	}
	return n*n + sumOfSquares(n-1)
}
