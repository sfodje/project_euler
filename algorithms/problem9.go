package algorithms

import "errors"

/*
a < b < c
(a*a) + (b*b) = (c*c)
a + b + c = n

math happens

a = ((2*b*n) - (n*n)) / ((2*b) - (2*n))

find the value of a for which a is a whole number
*/

func ProductOfSpecialPythTriplet(n int) float64 {
	for b := 1; b < n; b++ {
		if a, err := valueOfA(b, n); err == nil && isWholeNumber(a) {
			return a * float64(b) * (1000 - a - float64(b))
		}
	}
	return 0
}

func valueOfA(b, n int) (float64, error) {
	if b == n {
		return 0, errors.New("Error: b cannot be equal to n")
	}
	return float64(n*((2*b)-n)) / float64(2*(b-n)), nil
}

func isWholeNumber(v interface{}) bool {
	switch v := v.(type) {
	default:
		return false
	case int:
		return true
	case float32:
		return v == float32(int(v))
	case float64:
		return v == float64(int(v))
	}
}
