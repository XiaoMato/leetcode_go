package main

func getKthFromEnd(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}

	if k == 0 {
		return nil
	}

	tail := head
	result := head
	for ; k > 0; k-- {
		tail = tail.Next
	}

	for tail != nil {
		tail = tail.Next
		result = result.Next
	}
	return result
}
