package main

func rightSideView(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level > len(res) {
			res = append(res, root.Val)
		}
		//深度优先遍历的精髓就在下面两句中，根节点-》右子树-》左子树
		//注意是先访问右子树，所以每层第一个访问的节点永远是右子树节点
		//同理可以求解出左试图...
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(root, 1)
	return res
}
