package avlTree

import (
	"fmt"
)

type AvlNode struct {
	data     string
	priority int
}

func (node *AvlNode) compare(other *AvlNode) int {
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
	root   *AvlNode
	height int
	left   *AvlTree
	right  *AvlTree
}

//Creates an empty AVL tree
func NewAvlTree() *AvlTree {
	var tree AvlTree
	tree.height = -1
	return &tree
}

//Inserts node into *ptree and rebalances the tree
//Returns immediately if *ptree or node is nil
func Insert(ptree **AvlTree, node *AvlNode) {
	if ptree == nil || node == nil {
		return
	}
	tree := getTreePtrForInsert(ptree)
	if tree.root == nil {
		tree.root = node
		tree.height = 0
	} else {
		insertInChild(tree, node)
		updateHeightAndBalance(&tree)
	}
	*ptree = tree
}

//Returns max element in AVL tree
func Max(tree *AvlTree) *AvlNode {
	if tree == nil || tree.root == nil {
		return nil
	}
	if tree.right == nil {
		return tree.root
	}
	return Max(tree.right)
}

//Returns min element in AVL tree
func Min(tree *AvlTree) *AvlNode {
	if tree == nil || tree.root == nil {
		return nil
	}
	if tree.left == nil {
		return tree.root
	}
	return Min(tree.left)
}

//Returns true iff tree contains node
func Has(tree *AvlTree, node *AvlNode) bool {
	return findSubtreeWithNodeAsRoot(tree, node) != nil
}

func findSubtreeWithNodeAsRoot(tree *AvlTree, node *AvlNode) *AvlTree {
	if tree != nil && tree.root != nil && node != nil {
		rootToNodeCompare := tree.root.compare(node)
		if rootToNodeCompare == 0 {
			return tree
		}
		if rootToNodeCompare > 0 {
			return findSubtreeWithNodeAsRoot(tree.left, node)
		}
		return findSubtreeWithNodeAsRoot(tree.right, node)
	}
	return nil
}

func getTreePtrForInsert(ptree **AvlTree) *AvlTree {
	var tree *AvlTree
	if *ptree == nil {
		tree = NewAvlTree()
	} else {
		tree = *ptree
	}
	return tree
}

func insertInChild(tree *AvlTree, node *AvlNode) {
	if tree == nil {
		return
	}
	rootToNodeCompare := tree.root.compare(node)
	if rootToNodeCompare >= 0 {
		Insert(&tree.left, node)
	} else {
		Insert(&tree.right, node)
	}
}

func updateHeightAndBalance(ptree **AvlTree) {
	if ptree == nil || *ptree == nil {
		return
	}
	tree := *ptree
	prevHeight := tree.getHeight()
	tree.updateHeight()
	newHeight := tree.getHeight()
	if newHeight != prevHeight {
		balance(&tree)
	}
	*ptree = tree
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
	maxChildHeight := maxInt(leftHeight, rightHeight)
	return maxChildHeight + 1
}

func (tree *AvlTree) getHeight() int {
	if tree.isEmpty() {
		return -1
	}
	return tree.height
}

func balance(ptree **AvlTree) {
	if ptree == nil || *ptree == nil {
		return
	}
	tree := *ptree
	balance(&tree.left)
	balance(&tree.right)
	currBalance := tree.left.getHeight() - tree.right.getHeight()
	if currBalance > 1 {
		if tree.left.left.getHeight() > tree.left.right.getHeight() {
			rotateLeftToRoot(&tree)
		} else {
			doubleRotateLeftToRoot(&tree)
		}
	} else if currBalance < -1 {
		if tree.right.right.getHeight() > tree.right.left.getHeight() {
			rotateRightToRoot(&tree)
		} else {
			doubleRotateRightToRoot(&tree)
		}
	}
	*ptree = tree
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

//Graphical representation of doubleRotateLeftToRoot():
//          t                        t                     LR
//     L         R             LR         R           L          T
// LL    LR          ->     L     LRR         ->   LL  LRL    LRR  R
//     LRL LRR           LL  LRL
func doubleRotateLeftToRoot(ptree **AvlTree) {
	if ptree == nil || *ptree == nil {
		return
	}
	tree := *ptree
	if tree.left != nil && !tree.left.right.isEmpty() {
		rotateRightToRoot(&tree.left)
		rotateLeftToRoot(&tree)
		*ptree = tree
	}
}

func doubleRotateRightToRoot(ptree **AvlTree) {
	if ptree == nil || *ptree == nil {
		return
	}
	tree := *ptree
	if tree.right != nil && !tree.right.left.isEmpty() {
		rotateLeftToRoot(&tree.right)
		rotateRightToRoot(&tree)
		*ptree = tree
	}
}

func (tree *AvlTree) isEmpty() bool {
	return tree == nil || tree.root == nil
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
func maxInt(first int, second int) int {
	if first > second {
		return first
	}
	return second
}
