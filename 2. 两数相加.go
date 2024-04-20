package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret
	var plus int
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: (v1+v2+plus)%10}
		plus = (v1 + v2)/10
		cur = cur.Next
	}
	if plus != 0 {
		cur.Next = &ListNode{Val: plus}
	}
	return ret.Next
}
