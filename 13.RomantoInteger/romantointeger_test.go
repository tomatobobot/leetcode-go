package RomantoInteger

import "testing"

func TestRomanToInt(t *testing.T) {
	res := romanToInt("MCMXCIV")
	if res != 1994 {
		t.Fatal(res)
	}

	res = romanToInt("LVIII")
	if res != 58 {
		t.Fatal(res)
	}
	res = romanToInt("MMMMCM")
	if res != 4900 {
		t.Fatal(res)
	}
}
