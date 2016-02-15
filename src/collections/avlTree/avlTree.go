package avlTree

import (
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
	prevHeight := tree.getHeight()
	tree.updateHeight()
	newHeight := tree.getHeight()
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
	leftHeight := tree.left.getHeight()
	rightHeight := tree.right.getHeight()
	maxChildHeight := max(leftHeight, rightHeight)
	return maxChildHeight + 1
}

func (tree *AvlTree) getHeight() int {
	if tree.isEmpty() {
		return -1
	}
	return tree.height
}

//TODO
func (tree *AvlTree) balance() {
}

//Graphical representation of t.rotateLeft():
//       t                   tL
//   tL     tR     ->    tLL     t
//tLL tLR                      tLR tR
func (tree *AvlTree) rotateLeft() {
	if tree.isEmpty() {
		return
	}
	prevLeft := tree.left
	if prevLeft != nil {
		tree.left = prevLeft.right
		prevLeft.right = tree
		tree.updateHeight()
		prevLeft.updateHeight()
		tree = prevLeft
	}
}

//Graphical representation of t.rotateLeft():
//       t                   tR
//   tL     tR     ->     t     tRR
//        tRL tRR       tL tRL
func (tree *AvlTree) rotateRight() {
	if tree.isEmpty() {
		return
	}
	prevRight := tree.right
	if prevRight != nil {
		tree.right = prevRight.left
		prevRight.left = tree
		tree.updateHeight()
		prevRight.updateHeight()
		tree = prevRight
	}
}

//TODO: guard against tree.left == nil?
func (tree *AvlTree) doubleRotateLeft() {
	tree.left.rotateRight()
	tree.rotateLeft()
}

//TODO: guard against tree.right == nil?
func (tree *AvlTree) doubleRotateRight() {
	tree.right.rotateLeft()
	tree.rotateRight()
}

func (tree *AvlTree) isEmpty() bool {
	return tree == nil || tree.root == nil
}

func getTreeLeft(tree *AvlTree) *AvlTree {
	if tree == nil {
		return nil
	}
	return tree.left
}

func getTreeRight(tree *AvlTree) *AvlTree {
	if tree == nil {
		return nil
	}
	return tree.right
}

//Return max of two ints
func max(first int, second int) int {
	if first > second {
		return first
	}
	return second
}
