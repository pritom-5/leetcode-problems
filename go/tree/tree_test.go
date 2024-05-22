package tree

import "testing"

func createTree1 () *tree_t {
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

func TestDiameter(t *testing.T)  {
	tree_1 := createTree1()

	got := tree_1.root.diameter()
	if got != 3 {
		t.Errorf("got: %d, exp: %d", got, 3)
	}
}

func TestHeight (t *testing.T) {
	tree_1 := createTree1()
	height := tree_1.root.height()
	if height != 3 {
		t.Errorf("got: %d, exp: %d", height, 3)
	}
}

func TestIsBalanced (t *testing.T) {
	var got bool
	tree_1 := createTree1()

	
	if got = tree_1.root.isBalanced(); got != true {
		t.Errorf("got: %t, exp: true", got)
	}

	tree_2 := createTree2()

	if got = tree_2.root.isBalanced(); got != false {
		t.Errorf("got: %t, exp: false", got)
	}
}


func TestIsSameTree (t *testing.T) {
	tree1 := createTree1()
	tree2 := createTree1()

	

	got := isSameTree(tree1.root, tree2.root)
	if got != true {
		t.Errorf("got: %t, exp: true", got)
	}
}

func TestInvert (t *testing.T) {
	tree1 := createTree3()
	tree2 := createTree4()

	inverted_tree_root := invertTree(tree1.root)

	if isSameTree(tree2.root, inverted_tree_root) != true {
		t.Errorf("error invert")
	}
}





