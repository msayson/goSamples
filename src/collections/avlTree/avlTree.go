package avlTree

import (
	// "fmt" //XXX
	"sync"
)

type avlNode struct {
	data     string
	priority int
}

func (node *avlNode) compare(other *avlNode) int {
	if node == other {
		return 0
	}
	if other == nil || node.priority > other.priority {
		return 1
	}
	if node.priority < other.priority {
		return -1
	}
	if node.data > other.data {
		return 1
	}
	if node.data < other.data {
		return -1
	}
	return 0
}

type AvlTree struct {
	root   *avlNode
	height int
	left   *AvlTree
	right  *AvlTree
	lock   sync.RWMutex
}

func newAvlTree() *AvlTree {
	var tree AvlTree
	return &tree
}

func (tree *AvlTree) Insert(node *avlNode) {
	if node == nil {
		return
	}
	if tree == nil {
		tree = newAvlTree()
	}
	if tree.root == nil {
		tree.root = node
		tree.height = 0
	} else {
		tree.insertInChild(node)
		tree.updateHeightAndBalance()
	}
}

func (tree *AvlTree) insertInChild(node *avlNode) {
	rootToNodeCompare := tree.root.compare(node)
	if rootToNodeCompare >= 0 {
		tree.left.Insert(node)
	} else {
		tree.right.Insert(node)
	}
}

func (tree *AvlTree) updateHeightAndBalance() {
	prevHeight := getHeight(tree)
	tree.updateHeight()
	newHeight := getHeight(tree)
	if newHeight != prevHeight {
		tree.balance()
	}
}

//Updates height of tree based on heights of children
func (tree *AvlTree) updateHeight() {
	if !tree.isEmpty() {
		height := tree.calcHeightFromChildren()
		if tree.height != height {
			tree.height = height
		}
	}
}

func (tree *AvlTree) calcHeightFromChildren() int {
	if tree.isEmpty() {
		return -1
	}
	leftHeight := getHeight(tree.left)
	rightHeight := getHeight(tree.right)
	maxChildHeight := max(leftHeight, rightHeight)
	return maxChildHeight + 1
}

func getHeight(tree *AvlTree) int {
	if tree.isEmpty() {
		return -1
	}
	return tree.height
}

//TODO
func (tree *AvlTree) balance() {
}

func (tree *AvlTree) isEmpty() bool {
	return tree == nil || tree.root == nil
}

//Return max of two ints
func max(first int, second int) int {
	if first > second {
		return first
	}
	return second
}
