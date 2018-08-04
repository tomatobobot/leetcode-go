package ReverseInteger

import "testing"

func TestReverseInteger(t *testing.T) {
	x := reverseInteger(54321)
	if x != 12345 {
		t.Fatal(x)
	}
	x = reverseInteger(543210)
	if x != 12345 {
		t.Fatal(x)
	}
	x = reverseInteger(-54321)
	if x != -12345 {
		t.Fatal(x)
	}
	x = reverseInteger(1534236469)
	if x != 0 {
		t.Fatal(x)
	}
}
