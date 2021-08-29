package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || k == 0 {
		return head
	}

	var p = head
	length := 1
	for p.Next != nil {
		length++
		p = p.Next
	}
	k = length - k%length

	var result = head
	for i := 0; i < k; i++ {
		if result.Next == nil {
			result = head
		} else {
			result = result.Next
		}
	}

	var pointer = result
	for pointer.Next != result {
		if pointer.Next == nil {
			pointer.Next = head
		} else {
			pointer = pointer.Next
		}
	}
	pointer.Next = nil
	return result
}
