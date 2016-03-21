package algorithms

import "testing"

func TestLargestPalindromeProduct(t *testing.T) {
	testData := []struct {
		number        int
		expectedValue int
	}{
		{1, 9},
		{2, 9009},
		{3, 906609},
		{4, 99000099},
	}
	for _, test := range testData {
		if value := LargestPalidromeProduct(test.number); value != test.expectedValue {
			t.Errorf("LargestPalidromeProduct(%d) failed. Expected %d, Got %d",
				test.number, test.expectedValue, value)
		}
	}
}
