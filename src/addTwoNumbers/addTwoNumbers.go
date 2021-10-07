package main

import (
	"fmt"
)

/*
2. 两数相加 https://leetcode-cn.com/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*func initNode(nums []int) *ListNode{
	var head = new(ListNode)
	head.Val = nums[0]
	var p *ListNode = head
	for i := 1; i < len(nums); i++ {
		var node = new(ListNode)
		node.Val = nums[i]
		p.Next = node
		p = p.Next
	}
	return head
}*/
//切片转换为链表，漂亮
func initNode(nums []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	if nums != nil {
		for i := 0; i < len(nums); i++ {
			if head == nil {
				head = &ListNode{Val: nums[i]}
				tail = head
			} else {
				tail.Next = &ListNode{Val: nums[i]}
				tail = tail.Next
			}
		}
	}
	return head
}

func showNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d\t", l.Val)
		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	/*	//初始化链表一
		var  l1 = new(ListNode)
		l1.Val = 2
		var l1Node2 = new(ListNode)
		l1Node2.Val = 4
		l1.Next = l1Node2
		var l1Node3 = new(ListNode)
		l1Node3.Val = 3
		l1Node2.Next = l1Node3

		var  l2 = new(ListNode)
		l2.Val = 5
		var l2Node2 = new(ListNode)
		l2Node2.Val = 6
		l2.Next = l1Node2
		var l2Node3 = new(ListNode)
		l2Node3.Val = 4
		l2Node2.Next = l1Node3*/

	/*	//初始化链表二
		num1 := []int{2, 4, 3}
		var  l1 = new(ListNode)
		l1.Val = 2
		for i := 1; i < len(num1); i++ {
			var node = new(ListNode)
			node.Val = num1[i]
			l1.Next = node
		}
		num2 := []int{5, 6, 4}
		var  l2 = new(ListNode)
		l1.Val = 5
		for i := 1; i < len(num2); i++ {
			var node = new(ListNode)
			node.Val = num1[i]
			l2.Next = node
		}*/

	num1 := []int{2, 4, 3}
	num2 := []int{5, 6, 4}
	l1 := initNode(num1)
	l2 := initNode(num2)
	//showNode(l1)
	//showNode(l2)
	ret := addTwoNumbers(l1, l2)
	showNode(ret)
}
