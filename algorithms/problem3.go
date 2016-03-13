package algorithms

func LargestPrimeFactor(n int) int {
	if isPrime(n) {
		return n
	}
	for i, j := n/2, 2; i >= 2; i = n / j {
		if isAFactor(i, n) && isPrime(i) {
			return i
		}
		j += 1
	}
	return 0
}

func isAFactor(n, m int) bool {
	return m%n == 0
}

func isPrime(n int) bool {
	switch {
	case n < 2:
		return false
	case n <= 3 || n == 5 || n == 7:
		return true
	case n%2 == 0 || n%3 == 0 || n%5 == 0 || n%7 == 0:
		return false
	}
	for i, j := n/2, 2; i > 2; i = n / j {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		j += 1
	}
	return true
}
