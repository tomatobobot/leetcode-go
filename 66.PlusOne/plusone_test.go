package PlusOne

import "testing"

func TestPlusOne(t *testing.T) {
	digits := plusOne([]int{4, 9, 9, 9})
	if digits[0] != 5 {
		t.Fatal(digits)
	}
	digits = plusOne([]int{1, 2, 3})
	if digits[2] != 4 {
		t.Fatal(digits)
	}
	digits = plusOne([]int{9})
	if digits[0] != 1 {
		t.Fatal(digits)
	}
}
