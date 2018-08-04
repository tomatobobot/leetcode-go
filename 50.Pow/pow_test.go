package Pow

import "testing"

func TestMyPow(t *testing.T) {
	res := myPow(2.0, 10)
	if res != 1024.0 {
		t.Fatal(res)
	}
	res = myPow(2.0, -2)
	if res != 0.25 {
		t.Fatal(res)
	}
}
