package LongestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if res != "fl" {
		t.Fatal(res)
	}
	res = longestCommonPrefix([]string{"aa", "a"})
	if res != "a" {
		t.Fatal(res)
	}
}
