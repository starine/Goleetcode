package main

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil { //空结点找不到路径
		return false
	}
	//叶⼦结点，且路径和为sum
	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		return true
	}
	//递归进⼊⼦结点
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func pathSum1(root *TreeNode, target int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil { //空结点找不到路径
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 { //叶⼦结点，且路径和为target
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, target)
	return
}
