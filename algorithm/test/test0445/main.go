package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	count1 := count(l1)
	count2 := count(l2)
	if count1 > count2 {
		return merge(l1, l2, count1-count2)
	}
	return merge(l2, l1, count2-count1)
}

func count(l *ListNode) int {
	a := l
	num := 1
	for a.Next != nil {
		num++
		a = a.Next
	}
	return num
}

func merge(high *ListNode, low *ListNode, more int) *ListNode {
	count := 0
	for more != 0 {
		count = high.Val + count*10
		high = high.Next
		more--
	}
	for high != nil {
		high.Val = high.Val + low.Val
		high.Val = high.Val
		count = high.Val + count*10
		high = high.Next
		low = low.Next
	}
	l := &ListNode{}
	l.Next = nil
	if count == 0 {
		return l
	}
	for count != 0 {
		l.Val = count % 10
		count = count / 10
		a := &ListNode{}
		a.Next = l
		l = a
	}
	l = l.Next
	return l
}
