/*
* @Author: starine
* @Date:   2022/7/23 23:09
 */

package main

import "math"

//// 节点结构体
//type Node struct {
//	Index int
//	Val   int
//	Next  *Node
//}
//
//type NumberContainers struct {
//	DummyHead *Node //虚拟头节点
//	Size      int   //链表长度
//}
//
//func Constructor() NumberContainers {
//	dummyHead := &Node{
//		Index: 0,
//		Val:   -1,
//		Next:  nil,
//	}
//	return NumberContainers{dummyHead, 0}
//}
//
//func (this *NumberContainers) Change(index int, number int) {
//	cur := this.DummyHead
//	for cur != nil && cur.Next != nil && cur.Next.Index < index {
//		cur = cur.Next
//	}
//	if cur.Next != nil && index == cur.Next.Index {
//		cur.Next.Val = number
//	} else {
//		node := &Node{index, number, cur.Next}
//		cur.Next = node
//		this.Size++
//	}
//}
//
//func (this *NumberContainers) Find(number int) int {
//	cur := this.DummyHead
//	for cur != nil {
//		if cur.Val == number {
//			return cur.Index
//		} else {
//			cur = cur.Next
//		}
//	}
//	return -1
//}


type NumberContainers struct {
	mp map[int]int
	pos map[int]map[int]struct{}

}


func Constructor() NumberContainers {
	mp := make(map[int]int)
	pos := make(map[int]map[int]struct{})
	return NumberContainers{mp,pos}
}


func (this *NumberContainers) Change(index int, number int)  {
	if _,ok := this.mp[index];!ok {
		this.mp[index] = number
		m := map[int]struct{}{}
		m[index] = struct{}{}
		this.pos[number] =  m
	}else {
		m := this.pos[this.mp[index]]
		delete(m,index)
		this.mp[index] = number
		this.pos[number] = make(map[int]struct{})
		this.pos[number][index] = struct{}{}
	}

}


func (this *NumberContainers) Find(number int) int {
	m := this.pos[number]
	min := math.MaxInt
	for k, _ := range m {
		if k < min {
			min = k
		}
	}
	if min != math.MaxInt {
		return min
	}
	return -1
}
/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
