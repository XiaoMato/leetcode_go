package main


func splitListToParts(head *ListNode, k int) []*ListNode {
	var result []*ListNode
	p := head
	length := 0
	for p != nil {
		length++
		p = p.Next
	}
	n := length / k
	leftNum := length
	leftClm := k

	for head != nil {
		p = head
		var t int
		if leftNum%leftClm != 0 {
			t = n + 1
		} else {
			t = n
		}
		leftClm--
		leftNum = leftNum - t
		for i := 1; i < t; i++ {
			p = p.Next
		}
		result = append(result, head)
		if p == nil {
			break
		}
		head = p.Next
		p.Next = nil
	}
	for len(result) < k {
		result = append(result, nil)
	}
	return result
}
