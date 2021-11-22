package main

import "fmt"

/*
54. N 叉树的最大深度 leetcode559: https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
*/
type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}

func main() {
	//nums := []int{1,-1,3,2,4,-1,5,6} //N 叉树输入按层序遍历序列化表示，每组子节点由-1分隔
	var root Node
	root.Val = 1
	var node1, node2, node3, node4, node5 Node
	node1.Val = 3
	node2.Val = 2
	node3.Val = 4
	root.Children = append(root.Children, &node1)
	root.Children = append(root.Children, &node2)
	root.Children = append(root.Children, &node3)
	node4.Val = 5
	node5.Val = 6
	node3.Children = append(node3.Children, &node4)
	node3.Children = append(node3.Children, &node5)
	fmt.Println(maxDepth(&root))
}
