package leet

import "sort"

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

func at(node *TreeNode, index int) *TreeNode {
	if index < 0 {
		return nil
	}
	leftCount := 0
	if node.Left != nil {
		leftCount = count(node.Left)
	}
	if index < leftCount {
		return at(node.Left, index)
	}
	if index == leftCount {
		return node
	}

	if node.Right == nil {
		return nil
	}

	return at(node.Right, index - leftCount - 1)
}

func count(node *TreeNode) int {
	leftCount := 0
	if node.Left != nil {
		leftCount = count(node.Left)
	}
	rightCount := 0
	if node.Right != nil {
		rightCount = count(node.Right)
	}
	return leftCount + rightCount + 1
}

type treeNodesByVal struct {
	root *TreeNode
}

func (a *treeNodesByVal) Len() int           { return count(a.root) }

func (a *treeNodesByVal) Swap(i, j int) {
	iNode := at(a.root, i)
	jNode := at(a.root, j)

	iNode.Val, jNode.Val = jNode.Val, iNode.Val
}

func (a *treeNodesByVal) Less(i, j int) bool {
	iNode := at(a.root, i)
	jNode := at(a.root, j)

	return iNode.Val < jNode.Val
}

func RecoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	a := &treeNodesByVal{
		root:root,
	}
	sort.Sort(a)
}
