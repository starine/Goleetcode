package main

// queue
// 根据当前层的奇偶性输出vals，奇数层翻转vals即可
func levelOrder3(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	// 翻转数组
	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	level := 0             // 当前遍历到的层级
	q := []*TreeNode{root} // init & Push
	for len(q) > 0 {
		n, vals := len(q), []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:] // Pop
			vals = append(vals, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if level&1 == 1 { // 奇数层需要翻转
			reverse(vals)
		}
		res = append(res, vals)
		level++ // 遍历下一层
	}
	return res
}
