package main

import "fmt"

/*
16. 二叉搜索树中第K小的元素 leetcode230: https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

//先根遍历初始化二叉树
var i = -1

func HaveEmptyCreate(arr []int) *TreeNode {
	i = i + 1
	if i >= len(arr) {
		return nil
	}
	var t TreeNode
	if arr[i] != -1 {
		t = TreeNode{arr[i], nil, nil}
		t.Left = HaveEmptyCreate(arr)
		t.Right = HaveEmptyCreate(arr)
	} else {
		return nil
	}
	return &t
}

func main() {
	arr := []int{3, 1, -1, 2, 4} //-1表示nil
	root := HaveEmptyCreate(arr)
	k := 1
	fmt.Println(kthSmallest(root, k))
}
