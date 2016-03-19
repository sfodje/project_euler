package algorithms

import "math"

func SumOfPrimes(n int) int {
	return sumOfArray(getPrimes(n - 1))
}

func sumOfArray(arr *[]int) int {
	sum := 0
	for _, x := range *arr {
		sum += x
	}

	return sum
}

// sieve of Atkin algorithm
func getPrimes(limit int) *[]int {
	if limit < 2 {
		return &[]int{}
	}
	if limit == 2 {
		return &[]int{2}
	}
	if limit == 3 {
		return &[]int{2, 3}
	}
	sieve := make([]bool, limit+1)

	testingLimit := math.Ceil(math.Sqrt(float64(limit)))

	for i := float64(0); i < testingLimit; i++ {
		for j := float64(0); j < testingLimit; j++ {

			n := 4*int(math.Pow(i, 2)) + int(math.Pow(j, 2))
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				sieve[n] = !sieve[n]
			}

			n = int(3*math.Pow(i, 2)) + int(math.Pow(j, 2))
			if n <= limit && n%12 == 7 {
				sieve[n] = !sieve[n]
			}

			n = int(3*math.Pow(i, 2)) - int(math.Pow(j, 2))
			if n <= limit && i > j && n%12 == 11 {
				sieve[n] = !sieve[n]
			}
		}
	}

	for i := 5; i < int(testingLimit); i++ {
		if sieve[i] == true {
			k := int(math.Pow(float64(i), 2))
			for j := k; j <= limit; j += k {
				sieve[j] = false
			}
		}
	}

	primes := []int{2, 3}
	for i, isprime := range sieve {
		if isprime && i >= 5 {
			primes = append(primes, i)
		}
	}
	return &primes
}
