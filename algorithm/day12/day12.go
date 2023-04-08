package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) {

}
