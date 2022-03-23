package main

var stack1 []int //go使用数组实现栈（数组末尾即为栈顶）
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	for len(stack1) > 0 { //将第一个栈的元素弹出,并依次进入第二个栈中
		pop1 := stack1[len(stack1)-1]
		stack2 = append(stack2, pop1)
		stack1 = stack1[:len(stack1)-1]
	}
	pop := stack2[len(stack2)-1] //弹出第二个栈的顶部元素（数组末尾）
	stack2 = stack2[:len(stack2)-1]
	for len(stack2) > 0 {
		pop2 := stack2[len(stack2)-1]
		stack1 = append(stack1, pop2)
		stack2 = stack2[:len(stack2)-1]
	}
	return pop
}

func main() {

}
