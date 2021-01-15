package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var last, now *ListNode
	isFirst := true
	for {
		if isFirst {
			now = head.Next
			last = head
			last.Next = nil
			isFirst = false
			continue
		}
		if now.Next == nil {
			now.Next = last
			break
		}
		tmp := now
		now = now.Next
		tmp.Next = last
		last = tmp
	}
	return now
}
