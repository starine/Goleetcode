package main

func isCompleteTree(root *TreeNode) bool {
	if root == nil { //空树⼀定是完全⼆叉树
		return true
	}

	q := []*TreeNode{root}
	notComplete := false //定义⼀个⾸次出现nil的标记位
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == nil { //标记第⼀次遇到空节点
			notComplete = true
			continue
		}
		//队列不为空时，后续访问已经遇到空节点了，说明经过了叶⼦
		if notComplete {
			return false
		}
		q = append(q, cur.Left)
		q = append(q, cur.Right)
	}

	return true
}
