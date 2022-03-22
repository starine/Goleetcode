package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ { //遍历k次到尾部,也是下一组的头部
		if tail == nil {
			return head //如果不⾜k到了链表尾，直接返回，不翻转
		}
		tail = tail.Next
	}
	var pre *ListNode //翻转时需要的前序和当前节点
	cur := head
	for cur != tail { //在到达当前段尾节点前
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head.Next = reverseKGroup(tail, k) //当前尾指向下⼀段要翻转的链表
	return pre
}

func main() {

}
