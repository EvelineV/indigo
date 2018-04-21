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

func (n *Node) findMaxChild(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMaxChild(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("A nil node cannot be replaced.")
	}
	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (n *Node) deleteNode(s string, parent *Node) error {
	if n == nil {
		return errors.New("a nil node cannot be deleted.")
	}
	switch {
	case s < n.Index:
		return n.Left.deleteNode(s, n)
	case s > n.Index:
		return n.Right.deleteNode(s, n)
	default:
		// s == n.Index: remove this node
		if n.Left == nil && n.Right == nil {
			// this is a leaf node, set parent's Left and Right pointers to nil
			n.replaceNode(parent, nil)
			return nil
		}
		if n.Left == nil {
			// node only has right child, replace the node with its child
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			// node only has left child, replace the node with its child
			n.replaceNode(parent, n.Left)
			return nil
		}
		// node has two children. Replace node with its maximum element in the left subtree
		replacement, rParent := n.Left.findMaxChild(n)
		n.Index = replacement.Index
		n.Data = replacement.Data
		return replacement.deleteNode(replacement.Index, rParent)
	}
}

func (n *Node) size() int64 {
	if n == nil {
		return 0
	}
	return n.Left.size() + n.Right.size() + 1
}

func (n *Node) height() int64 {
	if n == nil {
		return 0
	}
	heightLeftSubtree := n.Left.height()
	heightRightSubtree := n.Right.height()
	if heightLeftSubtree > heightRightSubtree {
		return heightLeftSubtree + 1
	}
	return heightRightSubtree + 1
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

func (t *Tree) Delete(s string) error {
	if t.Root == nil {
		return nil
	}
	return t.Root.deleteNode(s, nil)
}

func (t *Tree) Size() int64 {
	if t.Root == nil {
		return 0
	}
	return t.Root.size()
}

func (t *Tree) Height() int64 {
	if t.Root == nil {
		return 0
	}
	return t.Root.height()
}
