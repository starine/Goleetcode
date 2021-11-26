package main

import "fmt"

//定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//定义二叉树
type BinareSearchTree struct {
	Root *TreeNode
}

//中序遍历二叉树
func midOrderTraversal(root *TreeNode) (res []int) {
	var midOrder func(node *TreeNode)
	midOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		midOrder(node.Left)
		res = append(res, node.Val)
		midOrder(node.Right)
	}
	midOrder(root)
	return
}

//二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			}
			p = p.Right
		}
	}
	return root
}

//700. 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val == root.Val {
			return root
		}
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 1, 3}
	root := &TreeNode{4, nil, nil}
	for _, num := range nums {
		insertIntoBST(root, num)
	}
	fmt.Println(levelOrder(root))
	val := 2
	ret := searchBST(root, val)
	fmt.Println(levelOrder(ret))
}
