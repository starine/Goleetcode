package main

func addInList(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = ReverseChain(l1)
	l2 = ReverseChain(l2)
	cur1, cur2 := l1, l2
	var head *ListNode //定义一个链表指针，每次相加之后让指针往前移
	carry := 0
	for cur1 != nil || cur2 != nil || carry > 0 {
		sum := 0
		if cur1 != nil {
			sum += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}
		sum += carry
		node := &ListNode{Val: sum % 10}
		node.Next = head //在链表头部插入
		head = node
		carry = sum / 10
	}
	return head
}
func ReverseChain(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}
	var pre *ListNode
	for list != nil {
		next := list.Next
		list.Next = pre
		pre = list
		list = next
	}
	return pre
} //941ms 27016kb

//使用栈实现
func addInList1(head1 *ListNode, head2 *ListNode) *ListNode {
	var head *ListNode
	var s1, s2 []int
	for head1 != nil {
		s1 = append(s1, head1.Val)
		head1 = head1.Next
	}
	for head2 != nil {
		s2 = append(s2, head2.Val)
		head2 = head2.Next
	}
	carry := 0
	for len(s1) > 0 || len(s2) > 0 || carry > 0 {
		sum := 0
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		sum += carry
		node := &ListNode{Val: sum % 10}
		node.Next = head
		head = node
		carry = sum / 10
	}
	return head
} //1024ms 31748kb
