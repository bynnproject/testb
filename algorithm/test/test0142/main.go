package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p1 := head.Next
	p2 := head.Next.Next
	for p2 != nil && p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}
	if p2 == nil {
		return nil
	}
	p1 = head
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
