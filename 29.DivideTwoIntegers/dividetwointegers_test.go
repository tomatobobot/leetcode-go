package DivideTwoIntegers

import "testing"

func TestDivide(t *testing.T) {
	res := divide(10, 3)
	if res != 3 {
		t.Fatal(res)
	}
	res = divide(7, -3)
	if res != -2 {
		t.Fatal(res)
	}

}
