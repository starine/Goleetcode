/*
* @Author: starine
* @Date:   2022/7/22 00:20
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	var l2 ListNode
	var l3 *ListNode
	l4 := ListNode{Val: 0, Next: nil}
	l5 := &ListNode{Val: 0, Next: nil}
	l6 := &ListNode{}
	fmt.Printf("%#v\n", l1)
	fmt.Printf("%#v\n", l2)
	fmt.Printf("%#v\n", l3)
	fmt.Printf("%#v\n", l4)
	fmt.Printf("%#v\n", l5)
	fmt.Printf("%#v\n", l6)

	num1 := []int{2, 4, 3, 5, 7, 9}
	head := slice2List(num1)
	printList(head)
	fmt.Println()

	reverse := reverseList(head)
	printList(reverse)
	fmt.Println()

	remove := removeElements(reverse, 9)
	printList(remove)
	fmt.Println()
}

//切片转换为链表
func slice2List(arr []int) *ListNode {
	var head, tail *ListNode
	if len(arr) == 0 { //nums == nil
		return head
	}
	head = &ListNode{Val: arr[0]} //头指针
	tail = head                   //尾指针
	for i := 1; i < len(arr); i++ {
		tail.Next = &ListNode{Val: arr[i]}
		tail = tail.Next
	}
	return head
}

//依次输出链表元素
func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
}

//单链表反转,反转之后head变成了尾指针
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		// 巧妙的利用go特性来进行节点指针位置交换
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

//删除链表指定元素
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
