package main

/*
请实现 copyRandomList 函数，复制一个复杂链表。
在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
*/
//复杂链表的复制

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 深度优先遍历的栈空间本身就是一个链式结构
var set map[*Node]struct{}

// 递归函数中的传参问题
func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}
	if _, ok := set[head]; !ok {
		return head
	}
	ans := new(Node)
	ans.Val = head.Val
	set[ans] = struct{}{}
	ans.Next = copyRandomList(head.Next)
	ans.Next = copyRandomList(head.Random)
	return ans
}

/*
type RandomListNode struct {
    Label int
    Next *RandomListNode
    Random *RandomListNode
}
*/

/**
 *
 * @param pHead RandomListNode类
 * @return RandomListNode类
 */
func Clone(head *RandomListNode) *RandomListNode {
	//write your code here
	if head == nil {
		return nil
	}
	if _, ok := set[head]; !ok {
		return head
	}
	ans := new(Node)
	ans.Val = head.Val
	set[ans] = struct{}{}
	ans.Next = copyRandomList(head.Next)
	ans.Next = copyRandomList(head.Random)
	return ans
}
