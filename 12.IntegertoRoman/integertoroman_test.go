package IntegertoRoman

import "testing"

func TestIntToRoman(t *testing.T) {
	res := intToRoman(8)
	if res != "VIII" {
		t.Fatal(res)
	}
	res = intToRoman(1994)
	if res != "MCMXCIV" {
		t.Fatal(res)
	}

	res = intToRoman(2994)
	if res != "MMCMXCIV" {
		t.Fatal(res)
	}
	res = intToRoman(58)
	if res != "LVIII" {
		t.Fatal(res)
	}
}
