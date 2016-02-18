package avlTree

import (
	"fmt"
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

func (tree *AvlTree) balance() {
	if tree == nil {
		return
	}
	tree.left.balance()
	tree.right.balance()
	currBalance := tree.left.getHeight() - tree.right.getHeight()
	if currBalance > 1 {
		if tree.left.left.getHeight() > tree.left.right.getHeight() {
			rotateLeftToRoot(&tree)
		} else {
			tree.doubleRotateLeftToRoot()
		}
	} else if currBalance < -1 {
		if tree.right.right.getHeight() > tree.right.left.getHeight() {
			rotateRightToRoot(&tree)
		} else {
			tree.doubleRotateRightToRoot()
		}
	}
}

//Graphical representation of rotateLeftToRoot(&t):
//       t                   tL
//   tL     tR     ->    tLL     t
//tLL tLR                      tLR tR
func rotateLeftToRoot(ptree **AvlTree) {
	if ptree == nil || (*ptree).isEmpty() {
		return
	}
	tree := *ptree
	prevLeft := tree.left
	if prevLeft != nil {
		tree.left = prevLeft.right
		prevLeft.right = tree
		tree.updateHeight()
		prevLeft.updateHeight()
		tree = prevLeft
	}
	*ptree = tree
}

//Graphical representation of rotateRightToRoot(&t):
//       t                   tR
//   tL     tR     ->     t     tRR
//        tRL tRR       tL tRL
func rotateRightToRoot(ptree **AvlTree) {
	if ptree == nil || (*ptree).isEmpty() {
		return
	}
	tree := *ptree
	prevRight := tree.right
	if !prevRight.isEmpty() {
		tree.right = prevRight.left
		prevRight.left = tree
		tree.updateHeight()
		prevRight.updateHeight()
		tree = prevRight
	}
	*ptree = tree
}

//Graphical representation of t.doubleRotateLeftToRoot():
//          t                        t                     LR
//     L         R             LR         R           L          T
// LL    LR          ->     L     LRR         ->   LL  LRL    LRR  R
//     LRL LRR           LL  LRL
func (tree *AvlTree) doubleRotateLeftToRoot() {
	if tree != nil && tree.left != nil && !tree.left.right.isEmpty() {
		rotateRightToRoot(&tree.left)
		rotateLeftToRoot(&tree)
	}
}

func (tree *AvlTree) doubleRotateRightToRoot() {
	if tree != nil && tree.right != nil && tree.right.left != nil {
		rotateLeftToRoot(&tree.right)
		rotateRightToRoot(&tree)
	}
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

func debug_printTree(tree *AvlTree, prefix string) {
	if tree == nil {
		fmt.Println(prefix + ": AvlTree<nil>")
	} else if !tree.isEmpty() {
		fmt.Printf(prefix+": %v\n  root: %v\n", tree, tree.root)
		debug_printTree(tree.left, prefix+".left")
		debug_printTree(tree.right, prefix+".right")
	}
}

//Return max of two ints
func max(first int, second int) int {
	if first > second {
		return first
	}
	return second
}
