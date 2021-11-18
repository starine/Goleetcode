package main

import "fmt"

/*
563. 二叉树的坡度 https://leetcode-cn.com/problems/binary-tree-tilt/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		ans += abs(sumLeft - sumRight)
		return sumLeft + sumRight + node.Val
	}
	dfs(root)
	return
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// BFS的数组翻译为二叉树
func createTreeNode(nums []int) *TreeNode {
	var root *TreeNode
	var queue []*TreeNode
	for i := 0; i < len(nums); i += 2 {
		if i == 0 {
			root = &TreeNode{Val: nums[i]}
			queue = append(queue, root)
		}
		parentNode := queue[0]
		queue = queue[1:]

		// 添加左节点
		if i+1 < len(nums) && nums[i+1] != -1001 { //题目范围为[-1000，1000]
			parentNode.Left = &TreeNode{Val: nums[i+1]}
			queue = append(queue, parentNode.Left)
		}

		// 添加右节点
		if i+2 < len(nums) && nums[i+2] != -1001 {
			parentNode.Right = &TreeNode{Val: nums[i+2]}
			queue = append(queue, parentNode.Right)
		}
	}
	return root
}

// 按层遍历
func LevelOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 采用队列实现
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree) // queue push
	for len(queue) > 0 {
		tree = queue[0]
		fmt.Printf(" %d -> ", tree.Val)

		queue = queue[1:] // queue pop

		if tree.Left != nil {
			queue = append(queue, tree.Left)
		}
		if tree.Right != nil {
			queue = append(queue, tree.Right)
		}
	}
}

func main() {
	arr := []int{21, 7, 14, 1, 1, 2, 2, 3, 3}
	root := createTreeNode(arr)
	LevelOrderTraversal(root)
	fmt.Println("")
	fmt.Println(findTilt(root))
}
