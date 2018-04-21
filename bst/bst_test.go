package bst

import "testing"

func TestInsertNode(t *testing.T) {
	root := &Node{Index: "n", Data: "This is the root node"}
	tree := &Tree{Name: "Testing Tree", Root: root}
	tree.Insert("q", "second node, to the right of the root.")
	if root.Right.Index != "q" {
		t.Errorf("second node is not on the right of the root!")
	}
	tree.Insert("a", "third node, should be leftmost")
	if root.Left.Index != "a" {
		t.Errorf("a < n")
	}
	tree.Insert("p", "between n and q")
	if root.Right.Left.Index != "p" {
		t.Errorf("P should be between n and q!")
	}
}

func TestFindNode(t *testing.T) {
	root := &Node{Index: "5", Data: "root"}
	tree := &Tree{Name: "test finding tree", Root: root}
	tree.Insert("3", "3")
	tree.Insert("9", "9")
	tree.Insert("7", "Node with index 7")
	d, found := tree.Find("7")
	if d != "Node with index 7" || found != true {
		t.Errorf("Uh oh... Search does not work!")
	}
}

func TestDeleteLeafNode(t *testing.T) {
	root := &Node{Index: "n", Data: "root"}
	tree := &Tree{Name: "test deleting leaf nodes", Root: root}
	tree.Insert("a", "a")
	tree.Insert("b", "b")
	tree.Insert("p", "p")
	tree.Insert("q", "q")
	if tree.Root.Left.Index != "a" {
		t.Errorf("failed setting up node a")
	}
	if tree.Root.Right.Right.Index != "q" {
		t.Errorf("failed setting up node q")
	}
	if tree.Root.Left.Right.Index != "b" {
		t.Errorf("failed setting up node b")
	}
	// start deleting
	tree.Delete("q")
	if tree.Root.Right.Right != nil {
		t.Errorf("failed deleting node q")
	}
	q, found := tree.Find("q")
	if found || q != "" {
		t.Errorf("failed deleting node q")
	}
	tree.Delete("b")
	if tree.Root.Left.Right != nil {
		t.Errorf("failed deleting node b")
	}
	b, found2 := tree.Find("b")
	if found2 || b != "" {
		t.Errorf("failed deleting node b")
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	root := &Node{Index: "n", Data: "root"}
	tree := &Tree{Name: "test deleting half-leaf nodes", Root: root}
	tree.Insert("a", "a")
	tree.Insert("b", "b")
	tree.Insert("p", "p")
	tree.Insert("q", "q")
	// start deleting
	tree.Delete("p")
	p, found := tree.Find("p")
	if found || p != "" {
		t.Errorf("failed deleting node p")
	}
	if tree.Root.Right.Index != "q" {
		t.Errorf("failed deleting node p and replacing it with q")
	}
	tree.Delete("a")
	a, found := tree.Find("a")
	if found || a != "" {
		t.Errorf("failed deleting node a")
	}
	if tree.Root.Left.Index != "b" {
		t.Errorf("failed deleting node a and replacing it with b")
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	tree := &Tree{Name: "Test deleting inner nodes", Root: &Node{Index: "root", Data: "root"}}
	tree.Insert("n", "node to be deleted, left child of the root")
	tree.Insert("p", "right child of n")
	tree.Insert("k", "left child of n, replaces n")
	tree.Insert("s", "right (only) child of p")
	tree.Insert("h", "left (only) child of k")
	tree.Insert("j", "right child of h")
	tree.Insert("c", "left child of h")
	tree.Delete("n")
	n, found := tree.Find("n")
	if found || n != "" {
		t.Errorf("failed deleting node n")
	}
	if tree.Root.Left.Index != "k" {
		t.Errorf("failed deleting node n and replacing it with k")
	}
	if tree.Root.Left.Left.Index != "h" {
		t.Errorf("could not find node h")
	}
}

func TestSize(t *testing.T) {
	tree := &Tree{Name: "Tree to calculate size", Root: &Node{Index: "7", Data: "root"}}
	tree.Insert("5", "5")
	tree.Insert("6", "6")
	tree.Insert("2", "2")
	tree.Insert("9", "9")
	size := tree.Size()
	if size != 5 {
		t.Errorf("Tree size should be 5 but is %d", &size)
	}
}

func TestHeight(t *testing.T) {
	tree := &Tree{Name: "Tree to calculate height", Root: &Node{Index: "7", Data: "root"}}
	tree.Insert("5", "5")
	tree.Insert("6", "6")
	tree.Insert("2", "2")
	tree.Insert("9", "9")
	tree.Insert("1", "1")
	height := tree.Height()
	if height != 4 {
		t.Errorf("Tree height should be 4 but is %d", &height)
	}
}
