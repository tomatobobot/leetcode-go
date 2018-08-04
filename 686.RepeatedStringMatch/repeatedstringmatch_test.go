package RepeatedStringMatch

import "testing"

func TestRepeatedStringMatch(t *testing.T) {
	res := repeatedStringMatch("abcd", "cdabcdab")
	if res != 3 {
		t.Fatal(res)
	}
	res = repeatedStringMatch("a", "aa")
	if res != 2 {
		t.Fatal(res)
	}
	res = repeatedStringMatch("abcabcabcabc", "abac")
	if res != -1 {
		t.Fatal(res)
	}
}
