package main

import "fmt"

/*
删除链表中的节点 leetcode237: https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func showNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d\t", l.Val)
		l = l.Next
	}
}
func main() {
	var node1 = new(ListNode)
	node1.Val = 4
	var node2 = new(ListNode)
	node2.Val = 5
	var node3 = new(ListNode)
	node3.Val = 1
	var node4 = new(ListNode)
	node4.Val = 9
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	deleteNode(node2)
	showNode(node1)
}
