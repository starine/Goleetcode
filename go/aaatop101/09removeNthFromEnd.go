package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Next: head} //添加表头,处理删掉第一个元素时比较方便
	pre := res                   //前序节点
	slow := head                 //当前节点
	fast := head
	for i := 0; i < n; i++ { //快指针先⾏n步
		fast = fast.Next
	}
	//快慢指针同步，快指针先到底，慢指针指向倒数第k个
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = slow.Next //删除该位置的节点
	return res.Next      //返回去掉头
}

func main() {

}
