package RemoveDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	s := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	if s != 5 {
		t.Fatal(s)
	}
}
