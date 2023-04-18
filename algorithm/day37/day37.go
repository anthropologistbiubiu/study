package main

import "fmt"

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, k int) *ListNode {
	var pre *ListNode = head
	var next *ListNode = head
	for i := 0; i < k; i++ {
		next = next.next
	}
	for next.next != nil {
		pre = pre.next
		next = next.next
	}
	pre.next = pre.next.next
	return head
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

func NewListNode(x int) *ListNode {
	return &ListNode{
		val:  x,
		next: nil,
	}
}

type ListNode struct {
	val  int
	next *ListNode
}
type List struct {
	l *ListNode
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	l := &List{
		l: nil,
	}
	for _, v := range nums {
		l.ListTailInsert(v)
	}
	ret := removeNthFromEnd(l.l, 2)
	for ret != nil {
		fmt.Println(ret.val)
		ret = ret.next
	}
}
