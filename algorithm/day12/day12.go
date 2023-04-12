package main

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var ret *ListNode = nil
	var cur *ListNode = nil
	var carry int = 0
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil && l2 != nil {
			carry += l1.val + l2.val
		} else if l1 != nil {
			carry += l1.val
		} else if l2 != nil {
			carry += l2.val
		}
		//fmt.Println(carry)
		val := carry % 10
		carry /= 10
		tem := NewListNode(val)
		if ret == nil {
			ret = tem
			cur = tem
		} else {
			cur.next = tem
			cur = cur.next
		}
		if l1 != nil {
			l1 = l1.next
		} else {
			l1 = nil
		}
		if l2 != nil {
			l2 = l2.next
		} else {
			l2 = nil
		}
	}
	return ret
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

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func main() {
	list1 := new(List)
	for _, val := range []int{2, 4, 3, 8, 1, 2, 9} {
		list1.ListTailInsert(val)
	}
	list2 := new(List)
	for _, val := range []int{5, 6, 4} {
		list2.ListTailInsert(val)
	}
	list1.ListPrint()
	list2.ListPrint()
	ret := addTwoNumbers(list1.l, list2.l)
	for ret != nil {
		fmt.Printf("%d->", ret.val)
		ret = ret.next

	}
	fmt.Println()
}
