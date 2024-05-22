package tree

import "math"

func (root *tree_node_t) diameter () int {
	length := 0
	diameterWalk(root, &length)
	return length
}

// walk dfs for diameter
func dfsWalk (node *tree_node_t, length *int) int {
	if node == nil {
		return 0
	}

	left := dfsWalk(node.Left, length)
	right := dfsWalk(node.Right, length)
	*length = max(*length, left + right)

	return max(left, right) + 1
}

func diameterWalk(node *tree_node_t, dia *int) int {
	if node == nil {
		return 0
	}
	left := diameterWalk(node.Left, dia)
	right := diameterWalk(node.Right, dia)
	*dia = max(left + right, *dia)

	return max(left, right) + 1
}

func (root *tree_node_t) height () int {
	return dfsWalkHeight(root)
}

func dfsWalkHeight (node *tree_node_t) int {
	if node == nil {
		return 0
	}
	left := dfsWalkHeight(node.Left)
	right := dfsWalkHeight(node.Right)

	return max(left, right) + 1
}


func (root *tree_node_t) isBalanced () bool {
	is_balanced := true
	walkIsBalanced(root, &is_balanced)
	return is_balanced
}

func walkIsBalanced (node *tree_node_t, is_balanced *bool) int {
	if node == nil {
		return 0
	}
	left := walkIsBalanced(node.Left, is_balanced)
	right := walkIsBalanced(node.Right, is_balanced)

	if math.Abs(float64(left - right)) > 1 {
		*is_balanced = false
	}

	return max(left, right) + 1
}


