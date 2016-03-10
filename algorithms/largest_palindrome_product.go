package algorithms

import "math"

func LargestPalidromeProduct(n int) int64 {
	largestNumber := int64(math.Pow(10, float64(n)) - 1)
	leastNumber := int64(math.Pow(10, float64(n-1)))
	palindrome := int64(0)
	for i := largestNumber; i >= leastNumber; i-- {
		for j := largestNumber; j >= leastNumber; j-- {
			product := i * j
			if product > palindrome && isPalindrome(product) {
				palindrome = product
			}
		}
	}
	return palindrome
}

func isPalindrome(n int64) bool {
	return firstHalf(n) == lastHalf(n)
}

func lastHalf(n int64) (half int64) {
	numOfDigits := numOfDigits(n)
	endDigit := numOfDigits - (numOfDigits / 2)
	for i := numOfDigits; i >= endDigit; i-- {
		digit := nthDigit(n, i)
		order := float64(i - (numOfDigits / 2))
		if numOfDigits%2 != 0 {
			order -= 1
		}
		half += int64(digit) * int64(math.Pow(10, order))
	}
	return
}

func firstHalf(n int64) int64 {
	numOfDigits := numOfDigits(n)
	numOfDigitsInHalf := numOfDigits / 2
	if numOfDigits%2 != 0 {
		numOfDigitsInHalf += 1
	}

	return n / int64(math.Pow(10, float64(numOfDigitsInHalf)))
}

func nthDigit(n int64, i int) int {
	divisor := int64(math.Pow(10, float64(numOfDigits(n)-i)))
	remainder := n % divisor
	if remainder < (divisor / 10) {
		return 0
	}
	return int(firstDigit(remainder))
}

func firstDigit(n int64) int {
	numOfDigits := numOfDigits(n) - 1
	return int(n / int64(math.Pow(10, float64(numOfDigits))))
}

func numOfDigits(n int64) int {
	if n < 0 {
		n *= -1
	}
	if n < 10 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(n)) + 1))
}
