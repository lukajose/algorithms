package tree

import (
	"collections/interfaces"
	"fmt"
)

type BinaryTree struct {
	root *Node
}

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	data   interfaces.ComparableItem
}

func newNode(left *Node, right *Node, parent *Node, data interfaces.ComparableItem) *Node {
	return &Node{
		left,
		right,
		parent,
		data,
	}
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

func (tree *BinaryTree) put(node *Node, data interfaces.ComparableItem) *Node {
	if node == nil {
		return newNode(nil, nil, nil, data)
	}
	switch node.data.CompareTo(data) {
	case 0:
		return node
	case -1:
		node.right = tree.put(node.right, data)
		break
	case 1:
		node.left = tree.put(node.left, data)
		break
	default:
		panic(fmt.Errorf("CompareTo is returning an output that does not meet the requirements"))
	}
	return node
}

func (tree *BinaryTree) Put(data interfaces.ComparableItem) {
	if tree.root == nil {
		tree.root = newNode(nil, nil, nil, data)
	}
	tree.root = tree.put(tree.root, data)
}

func (tree *BinaryTree) inorder(node *Node) {
	if node == nil {
		return
	}
	tree.inorder(node.left)
	fmt.Printf("%v ", node.data.Print())
	tree.inorder(node.right)
}

func (tree *BinaryTree) InOrder() {
	fmt.Println(" ==== tree ==== ")
	tree.inorder(tree.root)
	fmt.Println("")
}

func (tree *BinaryTree) SubTreeSearch(node *Node, data interfaces.ComparableItem) bool {
	if node == nil {
		return false
	}

	if node.data.CompareTo(data) == 0 {
		return true
	}

	if node.data.CompareTo(data) == -1 {
		return tree.SubTreeSearch(node.right, data)
	}

	return tree.SubTreeSearch(node.left, data)
}

func (tree *BinaryTree) Search(data interfaces.ComparableItem) bool {
	if tree.root == nil {
		return false
	}
	if tree.root.data.CompareTo(data) == 0 {
		return true
	}
	return tree.SubTreeSearch(tree.root, data)
}
