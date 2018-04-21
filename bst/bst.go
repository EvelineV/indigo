package bst

import (
	"errors"
)

type Node struct {
	Index string
	Data  string
	Left  *Node
	Right *Node
}

type Tree struct {
	Name string
	Root *Node
}

func (n *Node) insert(index, data string) error {
	if n == nil {
		return errors.New("Cannot insert an index into a nil tree!")
	}
	switch {
	case index == n.Index:
		return errors.New("A node with this index already exists!")

	case index < n.Index:
		if n.Left == nil {
			n.Left = &Node{Index: index, Data: data}
			return nil
		}
		return n.Left.insert(index, data)

	case index > n.Index:
		if n.Right == nil {
			n.Right = &Node{Index: index, Data: data}
			return nil
		}
		return n.Right.insert(index, data)
	}
	return nil
}

func (n *Node) find(s string) (string, bool) {
	if n == nil {
		return "", false
	}
	switch {
	case s == n.Index:
		return n.Data, true

	case s < n.Index:
		return n.Left.find(s)

	default:
		return n.Right.find(s)
	}
}

func (t *Tree) Insert(index, data string) error {
	if t.Root == nil {
		t.Root = &Node{Index: index, Data: data}
		return nil
	}
	return t.Root.insert(index, data)
}

func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.find(s)
}
