package main

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	l1 := ListLength(pHead1)
	l2 := ListLength(pHead2)
	if l1 > l2 { //当链表1更⻓时，链表1指针先⾛p1-p2步
		n := l1 - l2
		for i := 0; i < n; i++ {
			pHead1 = pHead1.Next
		}
		//两个链表同时移动，直到有公共结点时停下
		for pHead1 != nil && pHead2 != nil && pHead1 != pHead2 {
			pHead1 = pHead1.Next
			pHead2 = pHead2.Next
		}
	} else { //反之，则链表2先⾏p2-p1步
		n := l2 - l1
		for i := 0; i < n; i++ {
			pHead2 = pHead2.Next
		}
		for pHead1 != nil && pHead2 != nil && pHead1 != pHead2 {
			pHead1 = pHead1.Next
			pHead2 = pHead2.Next
		}
	}
	return pHead1
}

//计算链表⻓度的函数
func ListLength(head *ListNode) int {
	p := head
	n := 0
	for p != nil {
		n++
		p = p.Next
	}
	return n
}
