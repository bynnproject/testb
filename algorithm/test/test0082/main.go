package main

import "fmt"

func main() {
	l := &ListNode{Val: 1}
	pre := l
	l.Next = &ListNode{Val: 2}
	l = l.Next
	l.Next = &ListNode{Val: 3}
	l = l.Next
	l.Next = &ListNode{Val: 3}
	l = l.Next
	l.Next = &ListNode{Val: 5}
	l = l.Next
	l.Next = &ListNode{Val: 4}
	l = l.Next
	l.Next = &ListNode{Val: 4}
	l = l.Next
	res := deleteDuplicates(pre)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	record := make(map[int]int)
	tmp := head
	for tmp != nil {
		record[tmp.Val] += 1
		tmp = tmp.Next
	}
	pre := &ListNode{}
	res := pre
	for {
		for {
			val, _ := record[head.Val]
			if val == 1 {
				break
			}
			head = head.Next
			if head == nil {
				break
			}
		}

		if head != nil {
			fmt.Println(head.Val)
			pre.Next = head
			head = head.Next
			pre = pre.Next
		}

		if head == nil {
			pre.Next = nil
			break
		}
	}
	return res.Next
}
