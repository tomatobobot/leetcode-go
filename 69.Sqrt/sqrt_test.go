package Sqrt

import "testing"

func TestMySqrt(t *testing.T) {
	x := mySqrt(8)
	if x != 2 {
		t.Fatal(x)
	}
	x = mySqrt(4)
	if x != 2 {
		t.Fatal(x)
	}
}
