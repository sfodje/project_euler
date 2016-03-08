package algorithms

func LargestPrimeFactor(n int64) int64 {
	divisor := int64(1)
	for i := n; i >= 2; i = n / divisor {
		if isAFactor(i, n) && isPrime(i) {
			return i
		}
		divisor += 1
	}
	return 0
}

func isAFactor(n, m int64) bool {
	return m%n == 0
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
