package algorithms

func SumOfEvenFibonacci(n int64) int64 {
	return sumOfEvenFibonacci(n, 1, 2)
}

func sumOfEvenFibonacci(limit, f1, f2 int64) int64 {
	if f2 > limit || f1 > f2 {
		return 0
	}
	if f2%2 != 0 {
		return sumOfEvenFibonacci(limit, f2, f1+f2)
	}
	return f2 + sumOfEvenFibonacci(limit, f2, f1+f2)
}
