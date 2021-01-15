package main

import (
	"fmt"
)

func main() {
	l1 := NewList([]int{1, 2, 4})
	l2 := NewList([]int{1, 3, 4})
	printNode(l1)
	printNode(l2)
	l := mergeTwoLists(l1, l2)
	printNode(l)
}

func printNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func NewList(args []int) *ListNode {
	l := &ListNode{}
	b := l
	a := l
	for _, arg := range args {
		l.Val = arg
		b = l
		l.Next = &ListNode{}
		l = l.Next
	}
	b.Next = nil
	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 4
// 1 3 4
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	a := l
	b := l
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l.Val = l2.Val
			l2 = l2.Next
			l.Next = &ListNode{}
			b = l
			l = l.Next
			continue
		}
		if l2 == nil {
			l.Val = l1.Val
			l1 = l1.Next
			l.Next = &ListNode{}
			b = l
			l = l.Next
			continue
		}
		if l1.Val < l2.Val {
			l.Val = l1.Val
			l1 = l1.Next
			l.Next = &ListNode{}
			b = l
			l = l.Next
			continue
		}
		l.Val = l2.Val
		l2 = l2.Next
		l.Next = &ListNode{}
		b = l
		l = l.Next
	}
	b.Next = nil
	return a
}
