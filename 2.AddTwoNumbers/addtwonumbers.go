package AddTwoNumbers

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	carry := 0
	current := res
	for l1 != nil || l2 != nil || carry != 0 {
		num1, num2 := 0, 0
		if l1 != nil {
			num1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			num2, l2 = l2.Val, l2.Next
		}
		num := num1 + num2 + carry
		carry = num / 10
		current.Next = &ListNode{Val: num % 10, Next: nil}
		current = current.Next
	}
	return res.Next
}
