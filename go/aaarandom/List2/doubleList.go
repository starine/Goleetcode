/*
* @Author: starine
* @Date:   2022/7/22 13:37
 */

package List2

// 节点结构体
type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

// 链表对象
type MyLinkedList struct {
	DummyHead *Node
	DummyTail *Node
	Size      int
}

// 创建一个有虚拟首尾节点的链表
func Constructor() MyLinkedList {
	dummyHead := &Node{-1, nil, nil}
	dummyTail := &Node{
		Val:  -1,
		Next: nil,
		Pre:  dummyHead,
	}
	dummyHead.Next = dummyTail //指向尾结点
	return MyLinkedList{dummyHead, dummyTail, 0}
}

// 获取index节点
func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	} else {
		p := this.DummyHead
		for i := 0; i <= index; i++ { //找到的是index节点
			p = p.Next
		}
		return p.Val
	}
}

// 在头结点前面添加一个节点成为新的头节点
func (this *MyLinkedList) AddAtHead(val int) {
	// 直接在dummyHead后面添加一个节点即可
	var newNode *Node = &Node{val, nil, nil}
	// 后插法添加节点
	newNode.Next = this.DummyHead.Next // newnode指向前后两个节点
	newNode.Pre = this.DummyHead
	this.DummyHead.Next.Pre = newNode // 前后两个节点指向newNode
	this.DummyHead.Next = newNode
	this.Size++
}

// 在尾结点后面添加一个节点
func (this *MyLinkedList) AddAtTail(val int) {
	// 在DummyTail前面添加一个节点
	var newNode *Node = &Node{val, nil, nil}
	// 前插法添加节点
	newNode.Next = this.DummyTail // newnode指向前后两个节点
	newNode.Pre = this.DummyTail.Pre
	this.DummyTail.Pre.Next = newNode // 前后两个节点指向newnode
	this.DummyTail.Pre = newNode
	this.Size++
}

// 在指定位置位置插入一个节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	} else if index < 0 {
		this.AddAtHead(val)
	} else {
		// 找到索引为index.Pre的节点
		p := this.DummyHead
		for i := 0; i < index; i++ {
			p = p.Next
		}
		// 在index节点处插入
		var newNode *Node = &Node{Val: val, Pre: p, Next: p.Next}
		p.Next.Pre = newNode
		p.Next = newNode
		this.Size++
	}
}

// 删除指定index的节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size || index < 0 {
		return
	} else {
		p := this.DummyHead
		for i := 0; i < index; i++ { //找到的是index-1节点
			p = p.Next
		}
		p.Next.Next.Pre = p
		p.Next = p.Next.Next
		this.Size--
	}
}
