// 编写一个程序，找到两个单链表相交的起始节点。

package main

import "fmt"

/*
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
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

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return &ListNode{}
	}
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var ret *ListNode
	var tail *ListNode
	cur1 := l1
	cur2 := l2
	for cur1 != nil && cur2 != nil {
		x := 0
		if cur1.val <= cur2.val {
			x = cur1.val
			cur1 = cur1.next
		} else {
			x = cur2.val
			cur2 = cur2.next
		}
		tem := NewListNode(x)
		if ret == nil {
			ret = tem
			tail = tem
		} else {
			tail.next = tem
			tail = tail.next
		}
	}
	if cur1 != nil {
		tail.next = cur1
	}
	if cur2 != nil {
		tail.next = cur2
	}
	return ret
}
func main() {

	l1 := &List{
		l: nil,
	}
	for _, x := range []int{1, 2, 4} {
		l1.ListTailInsert(x)
	}
	l2 := &List{
		l: nil,
	}
	for _, x := range []int{1, 3, 4} {
		l2.ListTailInsert(x)
	}
	ret := mergeTwoLists(l1.l, l2.l)
	for ret != nil {
		fmt.Printf("%d->", ret.val)
		ret = ret.next
	}
	fmt.Println()
}
