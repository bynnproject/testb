package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	final := l
	var pre *ListNode
	add := 0
	for l1 != nil || l2 != nil {
		pre = l
		if l1 != nil && l2 != nil {
			l.Val = l1.Val + l2.Val + add
			add = l.Val / 10
			l.Val = l.Val % 10
			l.Next = &ListNode{}
			l = l.Next
			l1 = l1.Next
			l2 = l2.Next
			continue
		}
		if l1 == nil {
			l.Val = l2.Val + add
			add = l.Val / 10
			l.Val = l.Val % 10
			l.Next = &ListNode{}
			l = l.Next
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			l.Val = l1.Val + add
			add = l.Val / 10
			l.Val = l.Val % 10
			l.Next = &ListNode{}
			l = l.Next
			l1 = l1.Next
		}
	}
	if add != 0 {
		l.Val = add
	} else if pre != nil {
		pre.Next = nil
	}
	return final
}
