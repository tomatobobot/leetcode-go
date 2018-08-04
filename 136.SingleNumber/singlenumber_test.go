package SingleNumber

import "testing"

func TestSingleNumber(t *testing.T) {
	result := singleNumber([]int{4, 3, 2, 3, 1, 1, 2})
	if result != 4 {
		t.Fatal(result)
	}
}
