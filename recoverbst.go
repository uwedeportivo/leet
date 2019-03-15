package leet

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

func findSwap(node *TreeNode, validMin, validMax int64, pa, pb **TreeNode) {
	found := false
	if node.Val < validMin || node.Val > validMax {
		found = true
		if *pa == nil {
			*pa = node
		} else {
			*pb = node
			return
		}
	}

	if node.Left != nil {
		rangeEdge := node.Val
		if found {
			rangeEdge = validMax
		}
		findSwap(node.Left, validMin, rangeEdge, pa, pb)
	}
	if node.Right != nil {
		rangeEdge := node.Val
		if found {
			rangeEdge = validMin
		}
		findSwap(node.Right, rangeEdge, validMax, pa, pb)
	}
}

func RecoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var nodeA, nodeB *TreeNode

	findSwap(root, -infinity, infinity, &nodeA, &nodeB)

	if nodeA == nil {
		nodeA = root
	}

	if nodeB == nil {
		nodeB = root
	}

	nodeA.Val, nodeB.Val = nodeB.Val, nodeA.Val
}
