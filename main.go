package main

import (
	"flag"
	"fmt"

	. "github.com/sfodje/project_euler/algorithms"
)

func main() {
	problem := flag.Int("problem", 0, "Problem number")
	end := flag.Int("end", 0, "Last problem in range of problems")
	flag.Parse()

	if *problem < 0 || *end < 0 || (*end > 0 && *problem > *end) {
		panic("Error! Please check parameters")
	}

	if *end == 0 && *problem > 0 {
		end = problem
	}

	for i := *problem; i <= *end; i++ {
		if i == 0 || i == 1 {
			fmt.Printf("Problem 1: Sum of multiples of 3 and 5 below 1000 = %d\n", SumOfMultiplesOf3And5(1000))
		}

		if i == 0 || i == 2 {
			fmt.Printf("Problem 2: Sum of even fibonacci numbers 4,000,000 and below = %d\n", SumOfEvenFibonacci(4000000))
		}

		if i == 0 || i == 3 {
			fmt.Printf("Problem 3: Largest prime factor of 600851475143 = %d\n", LargestPrimeFactor(600851475143))
		}
		if i == 0 || i == 4 {
			fmt.Printf("Problem 4: Largest palindrome from product of 2 3-digit numbers = %d\n", LargestPalidromeProduct(3))
		}
		if i == 0 || i == 5 {
			multiple, _ := SmallestMultiple(1, 20)
			fmt.Printf("Problem 5: Smallest multiple of all numbers (1 to 20) = %d\n", multiple)
		}
		if i == 0 || i == 6 {
			fmt.Printf("Problem 6: Difference between the square of sum and sum of squares (1 to 100) = %d\n", SumSquareDifference(100))
		}
		if i == 0 || i == 7 {
			fmt.Printf("Problem 7: 10001th prime number = %d\n", NthPrimeNumber(10001))
		}
		if i == 0 || i == 8 {
			fmt.Printf("Problem 8: Largest product of 13 adjacent digits in thousand digit number = %d\n", GreatestProductOfNDigits(13))
		}
		if i == 0 || i == 9 {
			fmt.Printf("Problem 9: Product of special pythagorean triplet = %d\n", int64(ProductOfSpecialPythTriplet(1000)))
		}
		if i == 0 || i == 10 {
			fmt.Printf("Problem 10: Sum of all primes under 2,000,000 = %d\n", SumOfPrimes(2000000))
		}
		if i == 0 || i == 11 {
			fmt.Printf("Problem 11: Largest product of 2 adjacent numbers in number grid = %d\n", GetMaxProduct(4))
		}
		if i == 0 || i == 12 {
			fmt.Printf("Problem 12: Triangle number with more than 500 divisors = %d\n", TriangleNumberWithMoreThanNDivisors(500))
		}
		if i == 0 || i == 13 {
			fmt.Printf("Problem 13: The first 10 digits of the one-hundred 50-digit numbers = %d\n", FirstNthDigitsOfSum(10))
		}
	}
}
