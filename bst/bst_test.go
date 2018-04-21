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
