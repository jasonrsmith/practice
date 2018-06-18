package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{l1.Val + l2.Val, nil}
	runner := head
	l1 = l1.Next
	l2 = l2.Next
	var carry int
	if runner.Val >= 10 {
		carry = 1
		runner.Val = runner.Val - 10
	}
	for l1 != nil || l2 != nil || carry != 0 {
		runner.Next = &ListNode{}
		runner = runner.Next
		if l1 != nil {
			runner.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			runner.Val += l2.Val
			l2 = l2.Next
		}
		if carry != 0 {
			runner.Val += carry
		}
		if runner.Val >= 10 {
			carry = 1
			runner.Val = runner.Val - 10
		} else {
			carry = 0
		}
	}
	return head
}

func main() {
}
