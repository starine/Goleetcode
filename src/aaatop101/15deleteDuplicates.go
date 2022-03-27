package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil { //空链表
		return nil
	}

	cur := head           //遍历指针
	for cur.Next != nil { //指针当前和下⼀节点不为空
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next //如果当前与下⼀位相等则删掉下⼀节点
		} else { //否则指针正常遍历
			cur = cur.Next
		}
	}

	return head
}
