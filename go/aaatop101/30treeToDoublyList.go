package main

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}
	var inorder func(node *TreeNode) //准备闭包
	var head *TreeNode
	var pre *TreeNode

	inorder = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		inorder(cur.Left)
		if pre == nil {
			head = cur
			pre = cur
		} else {
			pre.Right = cur
			cur.Left = pre
			pre = cur
		}
		inorder(cur.Right)
	}

	inorder(pRootOfTree)
	//将最左孩子与最右孩子相互指向建立连接
	pre.Right = head
	head.Left = pre
	return head
}

// 左根右 使用栈中序遍历二叉树
func treeToDoublyList(root *TreeNode) *TreeNode {
	// 前驱节点，当前头节点
	var pre, head *TreeNode
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 中序遍历每次都从当前节点的最左孩子开始处理
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈处理当前节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 首次遍历，前驱节点为空，则记录头节点
		if pre == nil {
			head = root
		} else {
			// 非首次遍历，前驱节点与刚出栈的根节点相互指向建立连接
			pre.Right = root
			root.Left = pre
		}
		// 前驱节点记录当前根节点，切换到右孩子进行处理
		pre = root
		root = root.Right
	}
	// 单向循环结束以后，需要将最左孩子与最右孩子相互指向建立连接
	if pre != nil && head != nil {
		pre.Right = head
		head.Left = pre
	}
	return head
}
