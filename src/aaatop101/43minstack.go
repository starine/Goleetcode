package main

//用两个切片实现包含min函数的栈
//var stack1 [] int
//var stack2 [] int
//
//func Push(node int) {
//	stack1=append(stack1,node)
//	if len(stack2)==0 {
//		stack2=append(stack2,node)
//		return
//	}
//	if node>stack2[len(stack2)-1]{
//		stack2=append(stack2,stack2[len(stack2)-1])
//	}else{
//		stack2=append(stack2,node)
//	}
//	return
//}
//
//func Pop() {
//	stack1=stack1[:len(stack1)-1]
//	stack2=stack2[:len(stack2)-1]
//}
//func Top() int {
//	return stack1[len(stack1)-1]
//}
//func Min() int {
//	return stack2[len(stack2)-1]
//}

// MinStack define
type MinStack struct {
	elements, min []int
	l             int
}

// Constructor43 define
func Constructor43() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

// Push define
func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if this.l == 0 {
		this.min = append(this.min, x)
	} else {
		min := this.GetMin()
		if x < min {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.l++
}

func (this *MinStack) Pop() {
	this.l--
	this.min = this.min[:this.l]
	this.elements = this.elements[:this.l]
}

func (this *MinStack) Top() int {
	return this.elements[this.l-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.l-1]
}
