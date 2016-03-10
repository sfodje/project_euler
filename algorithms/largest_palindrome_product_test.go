package algorithms

import "testing"

func TestLargestPalindromeProduct(t *testing.T) {
	testData := []struct {
		number        int
		expectedValue int64
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

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		number        int64
		expectedValue bool
	}{
		{0, true},
		{1, true},
		{7, true},
		{10, false},
		{23, false},
		{305, false},
		{9009, true},
		{1015101, true},
		{101, true},
		{98766789, true},
		{102303201, true},
		{906609, true},
	}
	for _, test := range testData {
		if value := isPalindrome(test.number); value != test.expectedValue {
			t.Errorf("isPalindrome(%d) failed. Expected %d, Got %d", test.number, test.expectedValue, value)
		}
	}
}

func TestLastHalf(t *testing.T) {
	testData := []struct {
		number        int64
		expectedValue int64
	}{
		{0, 0},
		{1, 0},
		{7, 0},
		{10, 0},
		{23, 3},
		{305, 5},
		{357, 7},
		{3572, 27},
		{9009, 90},
	}
	for _, test := range testData {
		if value := lastHalf(test.number); value != test.expectedValue {
			t.Errorf("lastHalf(%d) failed. Expected %d, Got %d", test.number, test.expectedValue, value)
		}
	}
}

func TestFirstHalf(t *testing.T) {
	testData := []struct {
		number        int64
		expectedValue int64
	}{
		{0, 0},
		{1, 0},
		{7, 0},
		{10, 1},
		{23, 2},
		{305, 3},
		{357, 3},
		{3572, 35},
		{9009, 90},
	}
	for _, test := range testData {
		if value := firstHalf(test.number); value != test.expectedValue {
			t.Errorf("firstHalf(%d) failed. Expected %d, Got %d", test.number, test.expectedValue, value)
		}
	}
}

func TestNthDigit(t *testing.T) {
	testData := []struct {
		number        int64
		index         int
		expectedValue int
	}{
		{98765, 0, 9},
		{98765, 1, 8},
		{98765, 2, 7},
		{98765, 3, 6},
		{98765, 4, 5},
		{91036, 2, 0},
	}

	for _, test := range testData {
		if value := nthDigit(test.number, test.index); value != test.expectedValue {
			t.Errorf("nthDigit(%d, %d) failed. Expected %d, Got %d", test.number,
				test.index, test.expectedValue, value)
		}
	}
}

func TestFirstDigit(t *testing.T) {
	testData := []struct {
		number        int64
		expectedValue int
	}{
		{0, 0},
		{1, 1},
		{10, 1},
		{12345, 1},
		{54321, 5},
	}

	for _, test := range testData {
		if value := firstDigit(test.number); value != test.expectedValue {
			t.Errorf("firstDigit(%d) failed. Expected %d, Got %d", test.number, test.expectedValue, value)
		}
	}
}

func TestNumOfDigits(t *testing.T) {
	testData := make(map[int64]int)
	testData[0] = 1
	testData[1] = 1
	testData[10] = 2
	testData[89] = 2
	testData[-34] = 2
	testData[100] = 3
	testData[999] = 3
	testData[-1234567] = 7
	testData[123456789] = 9

	for arg, expectedValue := range testData {
		if value := numOfDigits(arg); value != expectedValue {
			t.Errorf("numOfDigits(%d) failed. Expected %d, got %d", arg, expectedValue, value)
		}
	}
}
