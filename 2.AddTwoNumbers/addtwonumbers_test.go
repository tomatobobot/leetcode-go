package AddTwoNumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	ln := addTwoNumbers(&ListNode{Val: 9}, &ListNode{Val: 9})
	if ln.Val != 8 {
		t.Fatal(ln)
	}
}
