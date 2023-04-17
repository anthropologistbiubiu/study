package main

/*
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
题目描述
编写一个程序，找到两个单链表相交的起始节点。
*/

/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func NewListNode(x int) *ListNode {
	return &ListNode{
		val:  x,
		next: nil,
	}
}
func removeElements(head *ListNode, val int) *ListNode {

	var pre *ListNode = head
	for pre.val == val && pre != nil {
		pre = pre.next
	}
	head = pre
	var cur *ListNode = pre.next
	for cur != nil {
		if cur.val == val {
			pre.next = cur.next
			cur = cur.next
		} else {
			pre = cur
			cur = cur.next
		}
	}
	return head
}

type List struct {
	l *ListNode
}

func (this *List) ListTailInsert(x int) {
	tem := NewListNode(x)
	cur := this.l
	if this.l == nil {
		this.l = tem
	} else {
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = tem
	}
}
func (this *List) ListPrint() {
	cur := this.l
	for cur != nil {
		fmt.Printf("%d->", cur.val)
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	list1 := new(List)
	for _, val := range []int{2, 4, 3, 8, 1, 2, 9} {
		list1.ListTailInsert(val)
	}
	ret := removeElements(list1.l, 2)
	for ret != nil {
		fmt.Printf("%d->", ret.val)
		ret = ret.next

	}
	fmt.Println()
}
