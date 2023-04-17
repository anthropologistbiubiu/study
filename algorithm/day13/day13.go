package main

import "fmt"

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

/*
题目描述
反转一个单链表。
示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

思路
这个就是常规操作了，使用一个变量记录前驱 pre，一个变量记录后继 next，不断更新current.next = pre 就好了。
链表的题目 90% 的 bug 都出现在：
头尾节点的处理
指针循环引用导致死循环
因此大家对这两个问题要保持 100% 的警惕。
关键点解析
链表的基本操作（交换）
虚拟节点 dummy 简化操作
注意更新 current 和 pre 的位置， 否则有可能出现溢出
*/
func reverseList(l1 *ListNode) *ListNode {
	cur := l1
	var pre *ListNode
	pre = nil
	asis := cur
	for cur != nil {
		asis = cur.next
		cur.next = pre
		pre = cur
		cur = asis

	}
	return pre
}

func main() {

	l1 := &List{
		l: nil,
	}
	for _, x := range []int{1, 2, 4} {
		l1.ListTailInsert(x)
	}
	ret := reverseList(l1.l)
	for ret != nil {
		fmt.Printf("%d->", ret.val)
		ret = ret.next
	}
	fmt.Println()
}
