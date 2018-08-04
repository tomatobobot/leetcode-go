package TwoSum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	ts := twoSum([]int{1, 2, 7, 11}, 9)
	if len(ts) != 2 || (ts[0] != 1 && ts[1] != 2) {
		t.Fatal(ts)
	}
}
