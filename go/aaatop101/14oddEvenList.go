package main

//双指针法
func oddEvenList(head *ListNode) *ListNode {
	if head == nil { //如果链表为空，不⽤重排
		return head
	}
	odd := head       //odd开头指向第⼀个节点
	even := head.Next //even开头指向第⼆个节点，可能为空
	evenHead := even  //指向even开头
	for even != nil && even.Next != nil {
		odd.Next = even.Next //odd连接even的后⼀个，即奇数位
		odd = odd.Next       //odd进⼊后⼀个奇数位
		even.Next = odd.Next //even连接后⼀个奇数的后⼀位，即偶数位
		even = even.Next     //even进⼊后⼀个偶数位
	}
	odd.Next = evenHead //even整体接在odd后⾯
	return head
}
