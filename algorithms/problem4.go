package algorithms

import (
	"math"
	"strconv"
)

func LargestPalidromeProduct(n int) int {
	largestNumber := int(math.Pow(10, float64(n)) - 1)
	leastNumber := int(math.Pow(10, float64(n-1)))
	palindrome := 0
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

func isPalindrome(n int) bool {
	return n == reverse(n)
}

func reverse(n int) int {
	s := strconv.Itoa(n)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	s = string(runes)
	n, _ = strconv.Atoi(s)
	return n
}
