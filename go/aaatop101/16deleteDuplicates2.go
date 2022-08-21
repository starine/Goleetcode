package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil { //空链表
		return nil
	}

	dummy := &ListNode{0, head} //在链表前加⼀个表头

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { //cur后面遇到相邻两个结点值相同
			tmp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next //将所有相同的都跳过
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next //返回时去掉增加的表头
}
