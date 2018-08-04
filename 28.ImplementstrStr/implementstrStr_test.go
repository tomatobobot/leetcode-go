package ImplementstrStr

import "testing"

func TestStrStr(t *testing.T) {
	s := strStr("hello", "")
	if s != 0 {
		t.Fatal(s)
	}
	s = strStr("hello", "ll")
	if s != 2 {
		t.Fatal(s)
	}

	s = strStr("hello", "hello world")
	if s != -1 {
		t.Fatal(s)
	}
}
