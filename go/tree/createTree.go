package tree

import "fmt"

type tree_node_t struct {
	Val int
	Left, Right *tree_node_t
}

type tree_t struct {
	root *tree_node_t
}

func (tree *tree_t) setRoot (val int) {
	tree.root = &tree_node_t{Val: val}
}
func (node *tree_node_t) insertLeft (val int) {
	node.Left = &tree_node_t{Val: val}
}
func (node *tree_node_t) insertRight (val int) {
	node.Right = &tree_node_t{Val: val}
}

func (tree *tree_t) printTree () {
	tree.root.printTree()
}

func (root *tree_node_t) printTree () {
	slice := make([]int, 0)
	q := make([]*tree_node_t, 0)
	q = append(q, root)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		slice = append(slice, curr.Val)

		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}

	}

	fmt.Printf("slice: %#v\n", slice)


}
