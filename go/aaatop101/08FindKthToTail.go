package main

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	fast := pHead
	slow := pHead
	for i := 0; i < k; i++ { //快指针先⾏k步
		if fast != nil {
			fast = fast.Next
		} else { //达不到k步说明链表过短，没有倒数k
			return nil
		}
	}
	//快慢指针同步，快指针先到底，慢指针指向倒数第k个
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {

}
