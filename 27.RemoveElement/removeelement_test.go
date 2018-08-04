package RemoveElement

import "testing"

func TestRemoveElement(t *testing.T) {
	count := removeElement([]int{1, 2, 3, 3, 2, 1}, 3)
	if count != 4 {
		t.Fatal(count)
	}
}
