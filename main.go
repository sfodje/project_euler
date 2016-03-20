package main

import (
	"fmt"

	. "github.com/sfodje/project_euler/algorithms"
)

func main() {
	fmt.Printf("Problem 1: Sum Multiples of 3 and 5 below 1000 = %d\n", SumOfMultiplesOf3And5(1000))
	fmt.Printf("Problem 2: Sum of Even Fibonacci Numbers <= 4,000,000 = %d\n", SumOfEvenFibonacci(4000000))
	fmt.Printf("Problem 3: Largest Prime Factor of 600851475143 = %d\n", LargestPrimeFactor(600851475143))
	fmt.Printf("Problem 4: Largest Palindrome From the Product of Two 3-Digit Numbers = %d\n", LargestPalidromeProduct(3))
	multiple, _ := SmallestMultiple(1, 20)
	fmt.Printf("Problem 5: Smallest Positive Number Evenly Divisible By All Numbers From 1 to 20 = %d\n", multiple)
	fmt.Printf("Problem 6: Difference Between Square of Sum And Sum Of Squares From 1 to 100 = %d\n", SumSquareDifference(100))
	fmt.Printf("Problem 7: The 10001th Prime Number = %d\n", NthPrimeNumber(10001))
	fmt.Printf("Problem 8: The Largest Product Of 13 Adjacent Digits in the Thousand Digit Number = %d\n", GreatestProductOfNDigits(13))
	fmt.Printf("Problem 9: The Product of the Special Pythagorean Triplet = %d\n", int64(ProductOfSpecialPythTriplet(1000)))
	fmt.Printf("Problem 10: The Sum of All Prime Numbers Under 2,000,000 = %d\n", SumOfPrimes(2000000))
	fmt.Printf("Problem 11: The Largest Product Of Four Adjacent Numbers In Grid = %d\n", GetMaxProduct(4))
}
