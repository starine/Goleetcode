package main

func hasCycle1(head *ListNode) *ListNode {
	if head == nil { //先判断链表为空的情况
		return nil
	}
	fast := head //快慢双指针
	slow := head
	for fast != nil && fast.Next != nil { //如果没环快指针会先到链表尾
		fast = fast.Next.Next //快指针移动两步
		slow = slow.Next      //慢指针移动⼀步
		if fast == slow {     //相遇则有环，返回相遇的位置
			return slow
		}
	}
	return nil //到末尾说明没有环，返回null
}
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow := hasCycle1(pHead)
	if slow == nil { //没有环
		return nil
	}
	fast := pHead      //快指针回到表头
	for fast != slow { //再次相遇即时是环入口
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {

}
