package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNodeš▒╗
 * @return ListNodeš▒╗
 */
func ReverseList(pHead *ListNode) *ListNode {
	var prev *ListNode
	for pHead != nil {
		next := pHead.Next
		pHead.Next = prev
		prev = pHead
		pHead = next
	}
	return prev
}

func main() {

}
