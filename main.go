package main

import (
	"fmt"

	"github.com/sfodje/project_euler/algorithms"
)

func main() {
	fmt.Printf("Problem 1: Sum Multiples of 3 and 5 below 1000 = %d\n", algorithms.SumOfMultiplesOf3And5(1000))
	fmt.Printf("Problem 2: Sum of Even Fibonacci Numbers <= 4,000,000 = %d\n", algorithms.SumOfEvenFibonacci(4000000))
	fmt.Printf("Problem 3: Largest Prime Factor of 600851475143 = %d\n", algorithms.LargestPrimeFactor(600851475143))
	fmt.Printf("Problem 4: Largest Palindrome From the Product of Two 3-Digit Numbers = %d\n", algorithms.LargestPalidromeProduct(3))
	multiple, _ := algorithms.SmallestMultiple(1, 20)
	fmt.Printf("Problem 5: Smallest Positive Number Evenly Divisible By All Numbers From 1 to 20 = %d\n", multiple)

	//fmt.Printf("Problem 548: Sum of Gozinta Chains For Which g(n) = n. Where n < 10^16 = %d\n", algorithms.GozintaProblem(10000000000000000))

}
