package PalindromeNumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	x := isPalindrome(121)
	if !x {
		t.Fatal(x)
	}
	x = isPalindrome(-121)
	if x {
		t.Fatal(x)
	}
}
