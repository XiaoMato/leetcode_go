package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}}}}}
	fmt.Printf("%+v", splitListToParts(head, 3))
}
