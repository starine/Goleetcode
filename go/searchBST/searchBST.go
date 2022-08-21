package main

import (
	"fmt"
	"math"
)

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

//使用栈实现中序遍历
func midOrderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
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

//判断是否是二叉搜索树
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

func (it *BSTIterator) Next() int {
	for node := it.cur; node != nil; node = node.Left {
		it.stack = append(it.stack, node)
	}
	it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
	val := it.cur.Val
	it.cur = it.cur.Right
	return val
}

func (it *BSTIterator) HasNext() bool {
	return it.cur != nil || len(it.stack) > 0
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		root = root.Right
		return root
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
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
	fmt.Println(isValidBST(root))
	fmt.Println(deleteNode(root, 3))
}
