package main

import "sort"

//链表归并排序 冒泡排序会超时
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	// 寻找中点并且要断开不然会影响后面寻找中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tmp := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(tmp)
	// 利用附加头节点方便返回头节点
	dummyNode := &ListNode{}
	cur := dummyNode
	// 具体归并过程
	for left != nil || right != nil {
		if left == nil {
			cur.Next = right
			right = right.Next
			cur = cur.Next
			continue
		}
		if right == nil {
			cur.Next = left
			left = left.Next
			cur = cur.Next
			continue
		}
		if left.Val <= right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	return dummyNode.Next
} //107ms 5344kb

//转换为数组进行排序
func sortInList(head *ListNode) *ListNode {
	var nums []int
	var p *ListNode = head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	sort.Ints(nums)
	p = head
	for _, num := range nums {
		p.Val = num
		p = p.Next
	}
	return head
} //116ms 5608kb
