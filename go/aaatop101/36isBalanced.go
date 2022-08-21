package main

//自底向上 边计算深度边判断是否平衡
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
