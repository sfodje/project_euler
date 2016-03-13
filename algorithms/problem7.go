package algorithms

func NthPrimeNumber(n int) int {
	num := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			num += 1
		}
		if num == n {
			return i
		}
	}
}
