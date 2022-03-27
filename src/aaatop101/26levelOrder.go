package main

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil { //如果是空，则直接返回空
		return ret
	}
	q := []*TreeNode{root} //队列存储，进⾏层次遍历
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{}) //记录⼆叉树的第i⾏
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ { //每层结点多少，队列⼤⼩就是多少
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			//若是左右孩⼦存在，则存⼊左右孩⼦作为下⼀个层次
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
