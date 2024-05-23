package tree

import (
	"math"
)

func (root *tree_node_t) diameter() int {
	length := 0
	diameterWalk(root, &length)
	return length
}

// walk dfs for diameter
func dfsWalk(node *tree_node_t, length *int) int {
	if node == nil {
		return 0
	}

	left := dfsWalk(node.Left, length)
	right := dfsWalk(node.Right, length)
	*length = max(*length, left+right)

	return max(left, right) + 1
}

func diameterWalk(node *tree_node_t, dia *int) int {
	if node == nil {
		return 0
	}
	left := diameterWalk(node.Left, dia)
	right := diameterWalk(node.Right, dia)
	*dia = max(left+right, *dia)

	return max(left, right) + 1
}

func (root *tree_node_t) height() int {
	return dfsWalkHeight(root)
}

func dfsWalkHeight(node *tree_node_t) int {
	if node == nil {
		return 0
	}
	left := dfsWalkHeight(node.Left)
	right := dfsWalkHeight(node.Right)

	return max(left, right) + 1
}

func (root *tree_node_t) isBalanced() bool {
	is_balanced := true
	walkIsBalanced(root, &is_balanced)
	return is_balanced
}

func walkIsBalanced(node *tree_node_t, is_balanced *bool) int {
	if node == nil {
		return 0
	}
	left := walkIsBalanced(node.Left, is_balanced)
	right := walkIsBalanced(node.Right, is_balanced)

	if math.Abs(float64(left-right)) > 1 {
		*is_balanced = false
	}

	return max(left, right) + 1
}

func isSameTree(p_node, q_node *tree_node_t) bool {
	if p_node == nil && q_node == nil {
		return true
	}
	if p_node == nil || q_node == nil || p_node.Val != q_node.Val {
		return false
	}

	return isSameTree(p_node.Left, q_node.Left) && isSameTree(p_node.Right, q_node.Right)
}

func invertTree(root *tree_node_t) *tree_node_t {
	if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	return root
}

func isSubtree(root, subtree *tree_node_t) bool {
	if root == nil {
		return false
	}

	return isEqual(root, subtree) || isSubtree(root.Left, subtree) || isSubtree(root.Right, subtree)

}

func isEqual(s, t *tree_node_t) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil || s.Val != t.Val {
		return false
	}

	return isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}

func hasPath(root *tree_node_t, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	tmp := targetSum - root.Val

	if hasPath(root.Left, tmp) || hasPath(root.Right, tmp) {
		return true
	}

	return false
}

func (root *tree_node_t) isBalanced02() bool {
	if root == nil {
		return false
	}
	l := isBalanced02_walk(root.Left)
	r := isBalanced02_walk(root.Right)

	return math.Abs(float64(l-r)) <= 1
}

func isBalanced02_walk(node *tree_node_t) int {
	if node == nil {
		return 0
	}
	l := isBalanced02_walk(node.Left)
	r := isBalanced02_walk(node.Right)
	return max(l, r) + 1

}

func (root *tree_node_t) isSymmetric() bool {
	
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return sym(root.Left, root.Right)
}

func sym(leftNode, rightNode *tree_node_t) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode == nil ||
		rightNode == nil ||
		leftNode.Val != rightNode.Val {
		return false
	}

	return sym(leftNode.Left, rightNode.Right) && sym(leftNode.Right, rightNode.Left)
}
