package algorithms

func SumOfMultiplesOf3And5(n int) int {
	n -= 1
	if n < 0 {
		return 0
	}
	if n%3 == 0 || n%5 == 0 {
		return n + SumOfMultiplesOf3And5(n)
	}
	return 0 + SumOfMultiplesOf3And5(n)
}

func SumOfMultiplesOf3And5Iterative(n int) int {
	if n < 3 {
		return 0
	}
	sum := 0
	for i := 3; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
