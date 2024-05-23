package tree

import (
	"testing"
)

func createTree1() *tree_t {
	tree := &tree_t{}
	tree.setRoot(1)
	tree.root.insertLeft(2)
	tree.root.insertRight(3)
	tree.root.Left.insertLeft(4)
	tree.root.Left.insertRight(5)

	return tree
}

func createTree2() *tree_t {
	tree := &tree_t{}
	tree.setRoot(1)
	tree.root.insertLeft(2)
	tree.root.insertRight(2)
	tree.root.Left.insertLeft(3)
	tree.root.Left.insertRight(3)
	tree.root.Left.Left.insertLeft(4)
	tree.root.Left.Left.insertRight(4)

	return tree
}

func createTree3() *tree_t {
	tree := &tree_t{}
	tree.setRoot(2)
	tree.root.insertLeft(1)
	tree.root.insertRight(3)

	return tree
}
func createTree4() *tree_t {
	tree := &tree_t{}
	tree.setRoot(2)
	tree.root.insertLeft(3)
	tree.root.insertRight(1)

	return tree
}

func createTree5() *tree_t {
	tree := &tree_t{}
	tree.setRoot(3)
	tree.root.insertLeft(4)
	tree.root.insertRight(5)
	tree.root.Left.insertLeft(1)
	tree.root.Left.insertRight(2)

	return tree
}

func createTree6() *tree_t {
	tree := &tree_t{}
	tree.setRoot(4)
	tree.root.insertLeft(1)
	tree.root.insertRight(2)

	return tree
}
func createTree7() *tree_t {
	tree := &tree_t{}
	tree.setRoot(1)

	return tree
}
func createTree8() *tree_t {
	tree := &tree_t{}
	tree.setRoot(0)

	return tree
}

func createTree9() *tree_t {
	tree := &tree_t{}
	tree.setRoot(5)
	tree.root.insertLeft(4)
	tree.root.Left.insertLeft(11)
	tree.root.Left.Left.insertLeft(7)
	tree.root.Left.Left.insertRight(2)

	tree.root.insertRight(8)
	tree.root.Right.insertLeft(13)
	tree.root.Right.insertRight(4)

	return tree
}

func createTree10() *tree_t {
	tree := &tree_t{}
	tree.setRoot(1)
	tree.root.insertLeft(2)
	tree.root.insertRight(2)
	tree.root.Left.insertLeft(3)
	tree.root.Right.insertRight(3)
	tree.root.Left.insertRight(4)
	tree.root.Right.insertLeft(4)

	return tree
}

func createTree11() *tree_t {
	tree := &tree_t{}
	tree.setRoot(1)
	tree.root.insertLeft(2)
	tree.root.insertRight(2)
	tree.root.Left.insertLeft(2)
	tree.root.Right.insertLeft(2)

	return tree
}

func TestDiameter(t *testing.T) {
	tree_1 := createTree1()

	got := tree_1.root.diameter()
	if got != 3 {
		t.Errorf("got: %d, exp: %d", got, 3)
	}
}

func TestHeight(t *testing.T) {
	tree_1 := createTree1()
	height := tree_1.root.height()
	if height != 3 {
		t.Errorf("got: %d, exp: %d", height, 3)
	}
}

func TestIsBalanced(t *testing.T) {
	var got bool
	tree_1 := createTree1()

	if got = tree_1.root.isBalanced02(); got != true {
		t.Errorf("got: %t, exp: true", got)
	}

	tree_2 := createTree2()

	if got = tree_2.root.isBalanced02(); got != false {
		t.Errorf("got: %t, exp: false", got)
	}
}

func TestIsSameTree(t *testing.T) {
	tree1 := createTree1()
	tree2 := createTree1()

	got := isSameTree(tree1.root, tree2.root)
	if got != true {
		t.Errorf("got: %t, exp: true", got)
	}
}

func TestInvert(t *testing.T) {
	tree1 := createTree3()
	tree2 := createTree4()

	inverted_tree_root := invertTree(tree1.root)

	if isSameTree(tree2.root, inverted_tree_root) != true {
		t.Errorf("error invert")
	}
}

func TestIsSubtree(t *testing.T) {
	tree1 := createTree5()
	tree2 := createTree6()

	got := isSubtree(tree1.root, tree2.root)

	if got != true {
		t.Errorf("got: %t, exp: %t", got, true)
	}

	tree3 := createTree7()
	tree4 := createTree8()

	if got = isSubtree(tree3.root, tree4.root); got != false {
		t.Errorf("got: %t, exp: %t", got, false)
	}
}

func TestHasPath(t *testing.T) {
	tree1 := createTree9()
	got := hasPath(tree1.root, 22)
	if got != true {
		t.Errorf("got: %t, exp: %t", got, true)
	}
}

func TestIsSymmetric(t *testing.T) {
	tree1 := createTree10()
	var got bool

	if got = tree1.root.isSymmetric(); got != true {
		t.Errorf("got: %t, exp: %t", got, true)
	}

	tree2 := createTree11()
	if got = tree2.root.isSymmetric(); got != false {
		t.Errorf("got: %t, exp: %t", got, false)
	}

}
