package avlTree

import (
	// "fmt"
	"runtime/debug"
	"testing"
)

func verifyCompareVal(t *testing.T, node *AvlNode, other *AvlNode, expectedVal int) {
	compareVal := node.compare(other)
	if compareVal != expectedVal {
		t.Errorf("%v.compare(%v) == %d, expected %d", *node, *other, compareVal, expectedVal)
		debug.PrintStack()
	}
}

func testNodeCompare_OtherIsNil(t *testing.T) {
	var nilNode AvlNode
	node := AvlNode{"A", 5}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &nilNode, expectedCompareResult)
}

func testNodeCompare_SameNode(t *testing.T) {
	node := AvlNode{"A", 5}
	expectedCompareResult := 0
	verifyCompareVal(t, &node, &node, expectedCompareResult)
}

func testNodeCompare_OtherIsEquivalent(t *testing.T) {
	node := AvlNode{"A", 5}
	equivNode := AvlNode{"A", 5}
	expectedCompareResult := 0
	verifyCompareVal(t, &node, &equivNode, expectedCompareResult)
}

func testNodeCompare_OtherHasLowerPriority(t *testing.T) {
	node := AvlNode{"A", 5}
	lowerPriorityNode := AvlNode{"B", 3}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &lowerPriorityNode, expectedCompareResult)
}

func testNodeCompare_OtherHasHigherPriority(t *testing.T) {
	node := AvlNode{"A", 5}
	higherPriorityNode := AvlNode{"B", 8}
	expectedCompareResult := -1
	verifyCompareVal(t, &node, &higherPriorityNode, expectedCompareResult)
}

func testNodeCompare_OtherHasSamePriorityLowerData(t *testing.T) {
	node := AvlNode{"ABC", 5}
	other := AvlNode{"AAA", 5}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &other, expectedCompareResult)
}

func testNodeCompare_OtherHasSamePriorityHigherData(t *testing.T) {
	node := AvlNode{"AAA", 5}
	other := AvlNode{"ABC", 5}
	expectedCompareResult := -1
	verifyCompareVal(t, &node, &other, expectedCompareResult)
}

func TestNodeCompare(t *testing.T) {
	testNodeCompare_OtherIsNil(t)
	testNodeCompare_SameNode(t)
	testNodeCompare_OtherIsEquivalent(t)
	testNodeCompare_OtherHasLowerPriority(t)
	testNodeCompare_OtherHasHigherPriority(t)
	testNodeCompare_OtherHasSamePriorityLowerData(t)
	testNodeCompare_OtherHasSamePriorityHigherData(t)
}

func TestNewAvlTree(t *testing.T) {
	tree := NewAvlTree()
	expectedHeight := -1
	if tree.height != expectedHeight {
		t.Errorf("tree.height == %d, expected %d", tree.height, expectedHeight)
	}
	if tree.root != nil {
		t.Errorf("tree.root == %v, expected nil", tree.root)
	}
	if tree.left != nil {
		t.Errorf("tree.left == %v, expected nil", tree.left)
	}
	if tree.right != nil {
		t.Errorf("tree.right == %v, expected nil", tree.right)
	}
}

func verifyTreeIsEmptyVal(t *testing.T, tree *AvlTree, expected bool) {
	isEmpty := tree.isEmpty()
	if isEmpty != expected {
		t.Errorf("isEmpty == %b, expected %b", isEmpty, expected)
		debug.PrintStack()
	}
}

func testTreeIsEmpty_EmptyTree(t *testing.T) {
	expected := true
	verifyTreeIsEmptyVal(t, nil, expected)

	var nilTree AvlTree
	verifyTreeIsEmptyVal(t, &nilTree, expected)

	emptyTree := NewAvlTree()
	verifyTreeIsEmptyVal(t, emptyTree, expected)
}

func testTreeIsEmpty_Leaf(t *testing.T) {
	expected := false
	leaf := createAvlTree_Leaf("a", 1)
	verifyTreeIsEmptyVal(t, leaf, expected)
}

func TestTreeIsEmpty(t *testing.T) {
	testTreeIsEmpty_EmptyTree(t)
	testTreeIsEmpty_Leaf(t)
}

func verifyTreeCalcHeightFromChildrenVal(t *testing.T, tree *AvlTree, expectedHeight int) {
	height := tree.calcHeightFromChildren()
	if height != expectedHeight {
		t.Errorf("calcHeightFromChildren() == %d, expected %d", height, expectedHeight)
		debug.PrintStack()
	}
}

func testTreeCalcHeightFromChildren_NilTree(t *testing.T) {
	expectedHeight := -1
	verifyTreeCalcHeightFromChildrenVal(t, nil, expectedHeight)

	var nilTree AvlTree
	verifyTreeCalcHeightFromChildrenVal(t, &nilTree, expectedHeight)
}

func testTreeCalcHeightFromChildren_Leaf(t *testing.T) {
	leaf := createAvlTree_Leaf("alpha", 3)
	expectedHeight := 0
	verifyTreeCalcHeightFromChildrenVal(t, leaf, expectedHeight)
}

func testTreeCalcHeightFromChildren_Parent(t *testing.T) {
	left := createAvlTree_Leaf("lower", 1)
	tree := createAvlTreeWithHeight("higher", 3, 1, left, nil)

	expectedHeight := 1
	verifyTreeCalcHeightFromChildrenVal(t, tree, expectedHeight)
}

func testTreeCalcHeightFromChildren_Grandparent(t *testing.T) {
	grandchild := createAvlTree_Leaf("low", 2)
	left := createAvlTreeWithHeight("lowest", 1, 1, nil, grandchild)
	tree := createAvlTreeWithHeight("high", 3, 2, left, nil)

	expectedHeight := 2
	verifyTreeCalcHeightFromChildrenVal(t, tree, expectedHeight)
}

func TestTreeCalcHeight(t *testing.T) {
	testTreeCalcHeightFromChildren_NilTree(t)
	testTreeCalcHeightFromChildren_Leaf(t)
	testTreeCalcHeightFromChildren_Parent(t)
	testTreeCalcHeightFromChildren_Grandparent(t)
}

func verifyGetHeightVal(t *testing.T, tree *AvlTree, expectedHeight int) {
	height := tree.getHeight()
	if height != expectedHeight {
		t.Errorf("getHeight() == %d, expected %d", height, expectedHeight)
		debug.PrintStack()
	}
}

func testTreeGetHeight_NilTree(t *testing.T) {
	expectedHeight := -1
	verifyGetHeightVal(t, nil, expectedHeight)

	var nilTree AvlTree
	verifyGetHeightVal(t, &nilTree, expectedHeight)
}

func testTreeGetHeight_Leaf(t *testing.T) {
	leafNode := createAvlNode("alpha", 3)
	var leaf AvlTree
	leafPtr := &leaf
	Insert(&leafPtr, leafNode)
	expectedHeight := 0
	verifyGetHeightVal(t, leafPtr, expectedHeight)
}

func TestTreeGetHeight(t *testing.T) {
	testTreeGetHeight_NilTree(t)
	testTreeGetHeight_Leaf(t)
}

func verifyUpdateHeight(t *testing.T, tree *AvlTree, expectedNewHeight int) {
	tree.updateHeight()
	updatedHeight := tree.getHeight()
	if updatedHeight != expectedNewHeight {
		t.Errorf("getHeight() == %d, expected %d", updatedHeight, expectedNewHeight)
		debug.PrintStack()
	}
}

func testTreeUpdateHeight_EmptyTree(t *testing.T) {
	expectedNewHeight := -1
	verifyUpdateHeight(t, nil, expectedNewHeight)

	var nilTree AvlTree
	verifyUpdateHeight(t, &nilTree, expectedNewHeight)

	emptyTree := NewAvlTree()
	verifyUpdateHeight(t, emptyTree, expectedNewHeight)
}

func testTreeUpdateHeight_Leaf(t *testing.T) {
	leaf := createAvlTree_Leaf("a", 1)

	expectedNewHeight := 0
	verifyUpdateHeight(t, leaf, expectedNewHeight)
}

func testTreeUpdateHeight_Parent(t *testing.T) {
	leaf := createAvlTree_Leaf("a", 1)
	parent := createAvlTreeWithHeight("b", 5, 1, leaf, nil)

	expectedNewHeight := 1
	verifyUpdateHeight(t, parent, expectedNewHeight)
}

func TestTreeUpdateHeight(t *testing.T) {
	testTreeUpdateHeight_EmptyTree(t)
	testTreeUpdateHeight_Leaf(t)
	testTreeUpdateHeight_Parent(t)
}

func verifyTreeRotateLeft_Empty(t *testing.T, tree *AvlTree) {
	if tree == nil {
		rotateLeftToRoot(&tree)
		verifyTreePointersEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		rotateLeftToRoot(&tree)
		verifyTreePointersEqual(t, tree, prevTree)
	}
}

func verifyTreeRotateRight_Empty(t *testing.T, tree *AvlTree) {
	if tree == nil {
		rotateRightToRoot(&tree)
		verifyTreePointersEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		rotateRightToRoot(&tree)
		verifyTreePointersEqual(t, tree, prevTree)
	}
}

func verifyTreePointersEqual(t *testing.T, tree *AvlTree, expected *AvlTree) {
	if tree != expected {
		if tree == nil {
			t.Errorf("tree == nil, expected %v\nexpected.root == %v", expected, expected.root)
		} else {
			t.Errorf("tree == %v, expected %v\ntree.root == %v", tree, expected, tree.root)
		}
		debug.PrintStack()
	}
}

func verifyTreeLAndR(t *testing.T, tree *AvlTree, expectedLeft *AvlTree, expectedRight *AvlTree) {
	verifyTreePointersEqual(t, tree.left, expectedLeft)
	verifyTreePointersEqual(t, tree.right, expectedRight)
}

func testTreeRotateLeft_EmptyTree(t *testing.T) {
	verifyTreeRotateLeft_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeRotateLeft_Empty(t, &nilTree)

	emptyTree := NewAvlTree()
	verifyTreeRotateLeft_Empty(t, emptyTree)
}

func testTreeRotateRight_EmptyTree(t *testing.T) {
	verifyTreeRotateRight_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeRotateRight_Empty(t, &nilTree)

	emptyTree := NewAvlTree()
	verifyTreeRotateRight_Empty(t, emptyTree)
}

func testTreeRotateLeft_Leaf(t *testing.T) {
	leaf := createAvlTree_Leaf("a", 1)
	prevLeafVal := *leaf
	rotateLeftToRoot(&leaf)
	if *leaf != prevLeafVal {
		t.Errorf("tree == &%v, expected &%v", *leaf, prevLeafVal)
	}
}

func testTreeRotateRight_Leaf(t *testing.T) {
	leaf := createAvlTree_Leaf("a", 1)
	prevLeafVal := *leaf
	rotateRightToRoot(&leaf)
	if *leaf != prevLeafVal {
		t.Errorf("tree == &%v, expected &%v", *leaf, prevLeafVal)
	}
}

func testTreeRotateLeft_ParentWithNoLeft(t *testing.T) {
	right := createAvlTree_Leaf("right", 2)
	parent := createAvlTreeWithHeight("parent", 1, 1, nil, right)
	prevParentVal := *parent
	rotateLeftToRoot(&parent)
	if *parent != prevParentVal {
		t.Errorf("tree == &%v, expected &%v", *parent, prevParentVal)
	}
}

func testTreeRotateLeft_ParentWithNoRight(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	tree := createAvlTreeWithHeight("parent", 1, 1, left, nil)
	prevTree := &*tree

	rotateLeftToRoot(&tree)
	if tree != left {
		t.Errorf("tree == &%v, expected &%v", *tree, *left)
	}
	verifyTreeLAndR(t, prevTree, nil, nil)
	verifyTreeLAndR(t, left, nil, prevTree)
	verifyGetHeightVal(t, left, 1)
	verifyGetHeightVal(t, prevTree, 0)
}

func testTreeRotateRight_ParentWithNoLeft(t *testing.T) {
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, nil, right)
	prevTree := &*tree

	rotateRightToRoot(&tree)
	if tree != right {
		t.Errorf("tree == %v, expected %v", tree, right)
	}
	verifyTreeLAndR(t, prevTree, nil, nil)
	verifyTreeLAndR(t, right, prevTree, nil)
	verifyGetHeightVal(t, right, 1)
	verifyGetHeightVal(t, prevTree, 0)
}

func testTreeRotateRight_ParentWithNoRight(t *testing.T) {
	left := createAvlTree_Leaf("left", 2)
	tree := createAvlTreeWithHeight("parent", 5, 1, left, nil)
	prevTree := &*tree
	rotateRightToRoot(&tree)
	if *tree != *prevTree {
		t.Errorf("tree == &%v, expected &%v", *tree, *prevTree)
	}
}

func testTreeRotateLeft_ParentWithLAndR(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevTree := &*tree
	prevLR := left.right

	rotateLeftToRoot(&tree)
	if tree != left {
		t.Errorf("New tree == %v, expected %v", tree, left)
	}
	verifyTreeLAndR(t, prevTree, prevLR, right)
	verifyTreeLAndR(t, left, nil, prevTree)
	verifyGetHeightVal(t, left, 2)
	verifyGetHeightVal(t, prevTree, 1)
}

func testTreeRotateRight_ParentWithLAndR(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevTree := &*tree
	prevRL := right.left

	rotateRightToRoot(&tree)
	if tree != right {
		t.Errorf("tree == %v, expected %v", tree, right)
	}
	verifyTreeLAndR(t, tree, prevTree, nil)
	verifyTreeLAndR(t, prevTree, left, prevRL)
	verifyGetHeightVal(t, right, 2)
	verifyGetHeightVal(t, prevTree, 1)
}

func testTreeRotateLeft_LongLeftTail(t *testing.T) {
	tL4 := createAvlTree_Leaf("LLLL", -25)
	tL3 := createAvlTreeWithHeight("LLL", -20, 1, tL4, nil)
	tL2 := createAvlTreeWithHeight("LL", -15, 2, tL3, nil)
	tLR := createAvlTree_Leaf("LR", -5)
	tL := createAvlTreeWithHeight("L", -10, 3, tL2, tLR)
	tR := createAvlTree_Leaf("R", 5)
	root := createAvlTreeWithHeight("root", 0, 4, tL, tR)
	prevTree := &*root

	rotateLeftToRoot(&root)
	if root != tL {
		t.Errorf("New tree == %v, expected %v", root, tL)
	}
	verifyTreeLAndR(t, tL, tL2, prevTree)
	verifyTreeLAndR(t, prevTree, tLR, tR)
	verifyGetHeightVal(t, tL, 3)
	verifyGetHeightVal(t, prevTree, 1)
}

func testTreeRotateRight_LongRightTail(t *testing.T) {
	tR4 := createAvlTree_Leaf("RRRR", 25)
	tR3 := createAvlTreeWithHeight("RRR", 20, 1, nil, tR4)
	tR2 := createAvlTreeWithHeight("RR", 15, 2, nil, tR3)
	tRL := createAvlTree_Leaf("R", 5)
	tR := createAvlTreeWithHeight("R", 10, 3, tRL, tR2)
	tL := createAvlTree_Leaf("L", -10)
	tree := createAvlTreeWithHeight("root", 0, 4, tL, tR)
	prevTree := &*tree

	rotateRightToRoot(&tree)
	if tree != tR {
		t.Errorf("New tree == %v, expected %v", tree, tR)
	}
	verifyTreeLAndR(t, tree, prevTree, tR2)
	verifyTreeLAndR(t, prevTree, tL, tRL)
	verifyGetHeightVal(t, tR, 3)
	verifyGetHeightVal(t, prevTree, 1)
}

func TestTreeRotateLeftToRoot(t *testing.T) {
	testTreeRotateLeft_EmptyTree(t)
	testTreeRotateLeft_Leaf(t)
	testTreeRotateLeft_ParentWithNoLeft(t)
	testTreeRotateLeft_ParentWithNoRight(t)
	testTreeRotateLeft_ParentWithLAndR(t)
	testTreeRotateLeft_LongLeftTail(t)
}

func TestTreeRotateRightToRoot(t *testing.T) {
	testTreeRotateRight_EmptyTree(t)
	testTreeRotateRight_Leaf(t)
	testTreeRotateRight_ParentWithNoLeft(t)
	testTreeRotateRight_ParentWithNoRight(t)
	testTreeRotateRight_ParentWithLAndR(t)
	testTreeRotateRight_LongRightTail(t)
}

func verifyTreeDoubleRotateLeft_Empty(t *testing.T, tree *AvlTree) {
	if tree == nil {
		doubleRotateLeftToRoot(&tree)
		verifyTreePointersEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		doubleRotateLeftToRoot(&tree)
		verifyTreePointersEqual(t, tree, prevTree)
	}
}

func verifyTreeDoubleRotateRight_Empty(t *testing.T, tree *AvlTree) {
	if tree == nil {
		doubleRotateRightToRoot(&tree)
		verifyTreePointersEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		doubleRotateRightToRoot(&tree)
		verifyTreePointersEqual(t, tree, prevTree)
	}
}

func testTreeDoubleRotateLeft_EmptyTree(t *testing.T) {
	verifyTreeDoubleRotateLeft_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeDoubleRotateLeft_Empty(t, &nilTree)

	emptyTree := NewAvlTree()
	verifyTreeDoubleRotateLeft_Empty(t, emptyTree)
}

func testTreeDoubleRotateRight_EmptyTree(t *testing.T) {
	verifyTreeDoubleRotateRight_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeDoubleRotateRight_Empty(t, &nilTree)

	emptyTree := NewAvlTree()
	verifyTreeDoubleRotateRight_Empty(t, emptyTree)
}

func testTreeDoubleRotateLeft_ParentWithNoLeft(t *testing.T) {
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, nil, right)
	prevTree := &*tree
	doubleRotateLeftToRoot(&tree)
	verifyTreePointersEqual(t, tree, prevTree)
}

func testTreeDoubleRotateRight_ParentWithNoRight(t *testing.T) {
	left := createAvlTree_Leaf("left", 2)
	tree := createAvlTreeWithHeight("parent", 5, 1, left, nil)
	prevTree := &*tree
	doubleRotateRightToRoot(&tree)
	verifyTreePointersEqual(t, tree, prevTree)
}

func testTreeDoubleRotateLeft_LeftWithNoRight(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevTree := &*tree
	doubleRotateLeftToRoot(&tree)
	verifyTreePointersEqual(t, tree, prevTree)
}

func testTreeDoubleRotateRight_RightWithNoLeft(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevTree := &*tree
	doubleRotateRightToRoot(&tree)
	verifyTreePointersEqual(t, tree, prevTree)
}

//    t              t              LR
// L     R   ->   LR    R   ->   L     t
//  LR           L                      R
func testTreeDoubleRotateLeft_LeftWithRight(t *testing.T) {
	leftR := createAvlTree_Leaf("leftR", -5)
	left := createAvlTreeWithHeight("left", -10, 1, nil, leftR)
	right := createAvlTree_Leaf("right", 5)
	tree := createAvlTreeWithHeight("parent", 0, 2, left, right)
	prevTree := &*tree

	doubleRotateLeftToRoot(&tree)
	if tree != leftR {
		debug_printTree(tree, "T")
		t.Errorf("tree == %v, expected %v", tree, leftR)
	}
	verifyTreeLAndR(t, tree, left, prevTree)
	verifyTreeLAndR(t, prevTree, nil, right)
	verifyTreeLAndR(t, left, nil, nil)
	verifyGetHeightVal(t, tree, 2)
	verifyGetHeightVal(t, prevTree, 1)
	verifyGetHeightVal(t, left, 0)
}

//    t              t              RL
// L     R   ->   LR   RL   ->   t     R
//     RL        L              L
func testTreeDoubleRotateRight_RightWithLeft(t *testing.T) {
	left := createAvlTree_Leaf("left", -5)
	rightL := createAvlTree_Leaf("rightL", 5)
	right := createAvlTreeWithHeight("right", 10, 1, rightL, nil)
	tree := createAvlTreeWithHeight("parent", 0, 2, left, right)
	prevTree := &*tree

	doubleRotateRightToRoot(&tree)
	if tree != rightL {
		t.Errorf("tree == %v, expected %v", tree, rightL)
	}
	verifyTreeLAndR(t, tree, prevTree, right)
	verifyTreeLAndR(t, prevTree, left, nil)
	verifyTreeLAndR(t, right, nil, nil)
	verifyTreeLAndR(t, left, nil, nil)
	verifyGetHeightVal(t, tree, 2)
	verifyGetHeightVal(t, prevTree, 1)
	verifyGetHeightVal(t, right, 0)
}

//          t                        t                     LR
//     L         R             LR         R           L          t
// LL    LR          ->     L     LRR         ->   LL  LRL    LRR  R
//     LRL LRR           LL  LRL
func testTreeDoubleRotateLeft_LeftWithGrandchildren(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -15)
	leftRL := createAvlTree_Leaf("LRL", -5)
	leftRR := createAvlTree_Leaf("LRR", -1)
	leftR := createAvlTreeWithHeight("LR", -3, 1, leftRL, leftRR)
	left := createAvlTreeWithHeight("L", -10, 2, leftL, leftR)
	right := createAvlTree_Leaf("R", 10)
	tree := createAvlTreeWithHeight("T", 5, 3, left, right)
	prevTree := &*tree

	doubleRotateLeftToRoot(&tree)
	if tree != leftR {
		t.Errorf("tree == %v, expected %v", tree, leftR)
	}
	verifyTreeLAndR(t, tree, left, prevTree)
	verifyTreeLAndR(t, prevTree, leftRR, right)
	verifyTreeLAndR(t, left, leftL, leftRL)
	verifyGetHeightVal(t, tree, 2)
	verifyGetHeightVal(t, prevTree, 1)
	verifyGetHeightVal(t, left, 1)
}

//      t                       t                        RL
// L         R             L         RL              t        R
//        RL   RR     ->          RLL    R     ->  L  RLL   RLR RR
//     RLL RLR                        RLR RR
func testTreeDoubleRotateRight_RightWithGrandchildren(t *testing.T) {
	left := createAvlTree_Leaf("L", -5)
	rightLL := createAvlTree_Leaf("RLL", 1)
	rightLR := createAvlTree_Leaf("RLR", 5)
	rightL := createAvlTreeWithHeight("RL", 3, 1, rightLL, rightLR)
	rightR := createAvlTree_Leaf("RR", 15)
	right := createAvlTreeWithHeight("R", 10, 2, rightL, rightR)
	tree := createAvlTreeWithHeight("T", 0, 3, left, right)
	prevTree := &*tree

	doubleRotateRightToRoot(&tree)
	if tree != rightL {
		t.Errorf("tree == %v, expected %v", tree, rightL)
	}
	verifyTreeLAndR(t, tree, prevTree, right)
	verifyTreeLAndR(t, prevTree, left, rightLL)
	verifyTreeLAndR(t, right, rightLR, rightR)
	verifyGetHeightVal(t, tree, 2)
	verifyGetHeightVal(t, prevTree, 1)
	verifyGetHeightVal(t, right, 1)
}

func TestTreeDoubleRotateLeftToRoot(t *testing.T) {
	testTreeDoubleRotateLeft_EmptyTree(t)
	testTreeDoubleRotateLeft_ParentWithNoLeft(t)
	testTreeDoubleRotateLeft_LeftWithNoRight(t)
	testTreeDoubleRotateLeft_LeftWithRight(t)
	testTreeDoubleRotateLeft_LeftWithGrandchildren(t)
}

func TestTreeDoubleRotateRightToRoot(t *testing.T) {
	testTreeDoubleRotateRight_EmptyTree(t)
	testTreeDoubleRotateRight_ParentWithNoRight(t)
	testTreeDoubleRotateRight_RightWithNoLeft(t)
	testTreeDoubleRotateRight_RightWithLeft(t)
	testTreeDoubleRotateRight_RightWithGrandchildren(t)
}

func verifyTreeBalanceHasNoEffect(t *testing.T, tree *AvlTree) {
	prevTree := &*tree
	balance(&tree)
	verifyTreePointersEqual(t, tree, prevTree)
}

func testTreeBalance_Nil(t *testing.T) {
	var tree *AvlTree = nil
	balance(&tree)
	verifyTreePointersEqual(t, tree, nil)
}

func testTreeBalance_AlreadyBalanced(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -5)
	verifyTreeBalanceHasNoEffect(t, leftL)

	leftR := createAvlTree_Leaf("LR", 5)
	left := createAvlTreeWithHeight("L", 1, 1, leftL, leftR)
	verifyTreeBalanceHasNoEffect(t, left)

	rightL := createAvlTree_Leaf("RL", 80)
	rightR := createAvlTree_Leaf("RR", 100)
	right := createAvlTreeWithHeight("R", 90, 1, rightL, rightR)
	tree := createAvlTreeWithHeight("root", 50, 2, left, right)
	verifyTreeBalanceHasNoEffect(t, tree)
}

func testTreeBalance_Empty(t *testing.T) {
	var nilTree AvlTree
	verifyTreeBalanceHasNoEffect(t, &nilTree)

	emptyTree := NewAvlTree()
	verifyTreeBalanceHasNoEffect(t, emptyTree)
}

func testTreeBalance_BalancesIn1LftRtt(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -3)
	left := createAvlTreeWithHeight("L", -1, 1, leftL, nil)
	tree := createAvlTreeWithHeight("root", 2, 2, left, nil)
	prevTree := &*tree

	balance(&tree)
	verifyTreePointersEqual(t, tree, left)
	verifyTreeLAndR(t, tree, leftL, prevTree)
	verifyTreeLAndR(t, prevTree, nil, nil)
	verifyTreeLAndR(t, leftL, nil, nil)
	verifyGetHeightVal(t, tree, 1)
	verifyGetHeightVal(t, prevTree, 0)
	verifyGetHeightVal(t, leftL, 0)
}

func testTreeBalance_BalancesIn1RghtRtt(t *testing.T) {
	rightR := createAvlTree_Leaf("RR", 3)
	right := createAvlTreeWithHeight("R", 1, 1, nil, rightR)
	tree := createAvlTreeWithHeight("root", -1, 2, nil, right)
	prevTree := &*tree

	balance(&tree)
	verifyTreePointersEqual(t, tree, right)
	verifyTreeLAndR(t, tree, prevTree, rightR)
	verifyTreeLAndR(t, prevTree, nil, nil)
	verifyTreeLAndR(t, rightR, nil, nil)
	verifyGetHeightVal(t, tree, 1)
	verifyGetHeightVal(t, prevTree, 0)
	verifyGetHeightVal(t, rightR, 0)
}

//          t                        t                     LR
//     L         R             LR         R           L          t
// LL    LR          ->     L     LRR         ->   LL  LRL    LRR  R
//     LRL LRR           LL  LRL
func testTreeBalance_BalancesIn1DblLftRtt(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -15)
	leftRL := createAvlTree_Leaf("LRL", -5)
	leftRR := createAvlTree_Leaf("LRR", -1)
	leftR := createAvlTreeWithHeight("LR", -3, 1, leftRL, leftRR)
	left := createAvlTreeWithHeight("L", -10, 2, leftL, leftR)
	right := createAvlTree_Leaf("R", 10)
	tree := createAvlTreeWithHeight("T", 5, 3, left, right)
	prevTree := &*tree

	balance(&tree)
	if tree != leftR {
		t.Errorf("tree == %v, expected %v", tree, leftR)
	}
	verifyTreeLAndR(t, tree, left, prevTree)
	verifyTreeLAndR(t, prevTree, leftRR, right)
	verifyTreeLAndR(t, left, leftL, leftRL)
	verifyGetHeightVal(t, tree, 2)
	verifyGetHeightVal(t, prevTree, 1)
	verifyGetHeightVal(t, left, 1)
}

//      t                       t                        RL
// L         R             L         RL              t        R
//        RL   RR     ->          RLL    R     ->  L  RLL   RLR RR
//     RLL RLR                        RLR RR
func testTreeBalance_BalancesIn1DblRghtRtt(t *testing.T) {
	left := createAvlTree_Leaf("L", -5)
	rightLL := createAvlTree_Leaf("RLL", 1)
	rightLR := createAvlTree_Leaf("RLR", 5)
	rightL := createAvlTreeWithHeight("RL", 3, 1, rightLL, rightLR)
	rightR := createAvlTree_Leaf("RR", 15)
	right := createAvlTreeWithHeight("R", 10, 2, rightL, rightR)
	tree := createAvlTreeWithHeight("T", 0, 3, left, right)
	prevTree := &*tree

	balance(&tree)
	if tree != rightL {
		t.Errorf("tree == %v, expected %v", tree, rightL)
	}
	verifyTreeLAndR(t, tree, prevTree, right)
	verifyTreeLAndR(t, prevTree, left, rightLL)
	verifyTreeLAndR(t, right, rightLR, rightR)
	verifyGetHeightVal(t, tree, 2)
	verifyGetHeightVal(t, prevTree, 1)
	verifyGetHeightVal(t, right, 1)
}

func TestTreeBalance(t *testing.T) {
	testTreeBalance_Nil(t)
	testTreeBalance_Empty(t)
	testTreeBalance_AlreadyBalanced(t)
	testTreeBalance_BalancesIn1LftRtt(t)
	testTreeBalance_BalancesIn1RghtRtt(t)
	testTreeBalance_BalancesIn1DblLftRtt(t)
	testTreeBalance_BalancesIn1DblRghtRtt(t)
}

func verifyNodePointersEqual(t *testing.T, node *AvlNode, expected *AvlNode) {
	if node != expected {
		t.Errorf("node == %v, expected %v", node, expected)
		debug.PrintStack()
	}
}

func testTreeInsert_NilTree(t *testing.T) {
	var nilTree *AvlTree = nil
	node := createAvlNode("a", 1)
	prevNode := &*node
	Insert(&nilTree, node)
	//Expect no effect, just verify no errors
	verifyCompareVal(t, node, prevNode, 0)
}

func testTreeInsert_EmptyTree(t *testing.T) {
	var tree AvlTree
	treePtr := &tree
	node := createAvlNode("a", 1)
	Insert(&treePtr, node)
	verifyNodePointersEqual(t, tree.root, node)
	verifyGetHeightVal(t, treePtr, 0)
}

func testTreeInsert_FirstChild_LowerPriority(t *testing.T) {
	tree := createAvlTree_Leaf("root", 2)
	newMinNode := createAvlNode("min", -1)
	Insert(&tree, newMinNode)

	if tree.left == nil {
		t.Errorf("tree.left == nil, expected subtree with node %v", newMinNode)
	} else {
		verifyNodePointersEqual(t, tree.left.root, newMinNode)
		verifyGetHeightVal(t, tree, 1)
	}
}

func testTreeInsert_FirstGrandchild_InitBalanced(t *testing.T) {
	left := createAvlTree_Leaf("left", 2)
	right := createAvlTree_Leaf("right", 8)
	tree := createAvlTreeWithHeight("root", 6, 1, left, right)

	grandchild := createAvlNode("newNode", 7)
	Insert(&tree, grandchild)
	if tree.right == nil {
		t.Errorf("tree.right == nil, expected non-nil tree")
	} else if tree.right.left == nil {
		t.Errorf("tree.right.left == nil, expected tree with node %v", grandchild)
		verifyNodePointersEqual(t, tree.right.left.root, grandchild)
		verifyGetHeightVal(t, tree, 2)
		verifyGetHeightVal(t, tree.left, 0)
		verifyGetHeightVal(t, tree.right, 1)
		verifyGetHeightVal(t, tree.right.left, 0)
	}
}

func testTreeInsert_LongTailShouldBalance(t *testing.T) {
	left := createAvlTree_Leaf("left", -5)
	tree := createAvlTreeWithHeight("root", -1, 1, left, nil)
	prevRoot := &*tree
	prevLeft := &*tree.left

	minNode := createAvlNode("min", -7)
	Insert(&tree, minNode)

	verifyNodePointersEqual(t, tree.root, prevLeft.root)
	verifyNodePointersEqual(t, tree.left.root, minNode)
	verifyNodePointersEqual(t, tree.right.root, prevRoot.root)
	verifyGetHeightVal(t, tree, 1)
	verifyGetHeightVal(t, tree.left, 0)
	verifyGetHeightVal(t, tree.right, 0)
}

func TestTreeInsert(t *testing.T) {
	testTreeInsert_NilTree(t)
	testTreeInsert_EmptyTree(t)
	testTreeInsert_FirstChild_LowerPriority(t)
	testTreeInsert_FirstGrandchild_InitBalanced(t)
	testTreeInsert_LongTailShouldBalance(t)
}

func testTreeMax_NilTree(t *testing.T) {
	var nilTree *AvlTree = nil
	max := Max(nilTree)
	verifyNodePointersEqual(t, max, nil)
}

func testTreeMax_EmptyTree(t *testing.T) {
	emptyTree := NewAvlTree()
	max := Max(emptyTree)
	verifyNodePointersEqual(t, max, nil)
}

func testTreeMax_SingleElement(t *testing.T) {
	node := createAvlNode("data", 6)
	tree := NewAvlTree()
	Insert(&tree, node)
	max := Max(tree)
	verifyNodePointersEqual(t, max, node)
}

func testTreeMax_NoRightNodes(t *testing.T) {
	tree := NewAvlTree()
	rootNode := createAvlNode("root", -6)
	leftNode := createAvlNode("left", -9)
	Insert(&tree, rootNode)
	Insert(&tree, leftNode)

	max := Max(tree)
	verifyNodePointersEqual(t, max, rootNode)
}

func testTreeMax_HasRightGrandchildOnLeft(t *testing.T) {
	left := createAvlTree_Leaf("L", 3)
	rightL := createAvlTree_Leaf("RL", 6)
	rightData := "R"
	rightPriority := 7
	right := createAvlTreeWithHeight(rightData, rightPriority, 1, rightL, nil)
	tree := createAvlTreeWithHeight("root", 5, 2, left, right)

	max := Max(tree)
	if max == nil {
		t.Errorf("Max(tree) == nil, expected &{%s %d}", rightData, rightPriority)
	} else if max.data != rightData || max.priority != rightPriority {
		t.Errorf("Max(tree) == %v, expected &{%s %d}", max, rightData, rightPriority)
	}
}

func testTreeMax_HasRightGrandchildOnRight(t *testing.T) {
	left := createAvlTree_Leaf("L", 3)
	rightL := createAvlTree_Leaf("RL", 6)
	maxData := "RR"
	maxPriority := 10
	rightR := createAvlTree_Leaf(maxData, maxPriority)
	right := createAvlTreeWithHeight("R", 8, 1, rightL, rightR)
	tree := createAvlTreeWithHeight("root", 5, 2, left, right)

	max := Max(tree)
	if max == nil {
		t.Errorf("Max(tree) == nil, expected &{%s %d}", maxData, maxPriority)
	} else if max.data != maxData || max.priority != maxPriority {
		t.Errorf("Max(tree) == %v, expected &{%s %d}", max, maxData, maxPriority)
	}
}

func TestTreeMax(t *testing.T) {
	testTreeMax_NilTree(t)
	testTreeMax_EmptyTree(t)
	testTreeMax_SingleElement(t)
	testTreeMax_NoRightNodes(t)
	testTreeMax_HasRightGrandchildOnLeft(t)
	testTreeMax_HasRightGrandchildOnRight(t)
}

func testTreeMin_NilTree(t *testing.T) {
	var nilTree *AvlTree = nil
	min := Min(nilTree)
	verifyNodePointersEqual(t, min, nil)
}

func testTreeMin_EmptyTree(t *testing.T) {
	emptyTree := NewAvlTree()
	min := Min(emptyTree)
	verifyNodePointersEqual(t, min, nil)
}

func testTreeMin_SingleElement(t *testing.T) {
	node := createAvlNode("data", 6)
	tree := NewAvlTree()
	Insert(&tree, node)
	min := Min(tree)
	verifyNodePointersEqual(t, min, node)
}

func testTreeMin_NoLeftNodes(t *testing.T) {
	tree := NewAvlTree()
	rootNode := createAvlNode("root", -6)
	rightNode := createAvlNode("right", 3)
	Insert(&tree, rootNode)
	Insert(&tree, rightNode)

	min := Min(tree)
	verifyNodePointersEqual(t, min, rootNode)
}

func testTreeMin_HasLeftGrandchildOnRight(t *testing.T) {
	leftR := createAvlTree_Leaf("LR", 4)
	leftData := "L"
	leftPriority := 3
	left := createAvlTreeWithHeight(leftData, leftPriority, 1, nil, leftR)
	right := createAvlTree_Leaf("R", 12)
	tree := createAvlTreeWithHeight("root", 5, 2, left, right)

	min := Min(tree)
	if min == nil {
		t.Errorf("Min(tree) == nil, expected &{%s %d}", leftData, leftPriority)
	} else if min.data != leftData || min.priority != leftPriority {
		t.Errorf("Min(tree) == %v, expected &{%s %d}", min, leftData, leftPriority)
	}
}

func testTreeMin_HasLeftGrandchildOnLeft(t *testing.T) {
	minData := "LL"
	minPriority := 2
	leftL := createAvlTree_Leaf(minData, minPriority)
	leftR := createAvlTree_Leaf("LR", 4)
	left := createAvlTreeWithHeight("L", 3, 1, leftL, leftR)

	right := createAvlTree_Leaf("R", 9)
	tree := createAvlTreeWithHeight("root", 5, 2, left, right)

	min := Min(tree)
	if min == nil {
		t.Errorf("Min(tree) == nil, expected &{%s %d}", minData, minPriority)
	} else if min.data != minData || min.priority != minPriority {
		t.Errorf("Min(tree) == %v, expected &{%s %d}", min, minData, minPriority)
	}
}

func TestTreeMin(t *testing.T) {
	testTreeMin_NilTree(t)
	testTreeMin_EmptyTree(t)
	testTreeMin_SingleElement(t)
	testTreeMin_NoLeftNodes(t)
	testTreeMin_HasLeftGrandchildOnRight(t)
	testTreeMin_HasLeftGrandchildOnLeft(t)
}

func testTreeHas_NilTree(t *testing.T) {
	var nilTree *AvlTree = nil
	node := createAvlNode("data", 5)
	hasNode := Has(nilTree, node)
	if hasNode {
		t.Errorf("Has(%v, %v) == true, expected false", nilTree, node)
	}
}

func testTreeHas_EmptyTree(t *testing.T) {
	var tree AvlTree
	node := createAvlNode("data", 5)
	hasNode := Has(&tree, node)
	if hasNode {
		t.Errorf("Has(&%v, %v) == true, expected false", tree, node)
	}
}

func testTreeHas_NilNode(t *testing.T) {
	tree := createAvlTree_Leaf("data", 5)
	var node *AvlNode = nil
	hasNode := Has(tree, node)
	if hasNode {
		t.Errorf("Has(%v, %v) == true, expected false", tree, node)
	}
}

func testTreeHas_IsRoot(t *testing.T) {
	node := createAvlNode("data", 6)
	tree := NewAvlTree()
	Insert(&tree, node)
	hasNode := Has(tree, node)
	if !hasNode {
		t.Errorf("Has(%v, %v) == false, expected true", tree, node)
	}
}

func testTreeHas_IsChild(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -1)
	leftR := createAvlTree_Leaf("LR", 3)
	data := "data"
	priority := 2
	left := createAvlTreeWithHeight(data, priority, 1, leftL, leftR)
	rightL := createAvlTree_Leaf("RL", 14)
	right := createAvlTreeWithHeight("R", 15, 1, rightL, nil)
	tree := createAvlTreeWithHeight("root", 10, 2, left, right)

	searchNode := createAvlNode(data, priority)
	hasNode := Has(tree, searchNode)
	if !hasNode {
		t.Errorf("Has(%v, %v) == false, expected true", tree, searchNode)
		debug_printTree(tree, "T")
	}
}

func testTreeHas_IsGrandchild(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -1)
	data := "data"
	priority := 3
	leftR := createAvlTree_Leaf(data, priority)
	left := createAvlTreeWithHeight("L", 2, 1, leftL, leftR)
	rightL := createAvlTree_Leaf("RL", 14)
	right := createAvlTreeWithHeight("R", 15, 1, rightL, nil)
	tree := createAvlTreeWithHeight("root", 10, 2, left, right)

	searchNode := createAvlNode(data, priority)
	hasNode := Has(tree, searchNode)
	if !hasNode {
		t.Errorf("Has(%v, %v) == false, expected true", tree, searchNode)
		debug_printTree(tree, "T")
	}
}

func testTreeHas_PriorityNotInTree(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -1)
	data := "data"
	leftR := createAvlTree_Leaf(data, 3)
	left := createAvlTreeWithHeight("L", 2, 1, leftL, leftR)
	rightL := createAvlTree_Leaf("RL", 14)
	right := createAvlTreeWithHeight("R", 15, 1, rightL, nil)
	tree := createAvlTreeWithHeight("root", 10, 2, left, right)

	unusedPriority := 4
	searchNode := createAvlNode(data, unusedPriority)
	hasNode := Has(tree, searchNode)
	if hasNode {
		t.Errorf("Has(%v, %v) == true, expected false", tree, searchNode)
		debug_printTree(tree, "T")
	}
}

func testTreeHas_DataNotInTree(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -1)
	leftR := createAvlTree_Leaf("LR", 3)
	left := createAvlTreeWithHeight("L", 2, 1, leftL, leftR)
	priority := 14
	rightL := createAvlTree_Leaf("RL", priority)
	right := createAvlTreeWithHeight("R", 15, 1, rightL, nil)
	tree := createAvlTreeWithHeight("root", 10, 2, left, right)

	unusedData := "data"
	searchNode := createAvlNode(unusedData, priority)
	hasNode := Has(tree, searchNode)
	if hasNode {
		t.Errorf("Has(%v, %v) == true, expected false", tree, searchNode)
		debug_printTree(tree, "T")
	}
}

func TestTreeHas(t *testing.T) {
	testTreeHas_NilTree(t)
	testTreeHas_EmptyTree(t)
	testTreeHas_NilNode(t)
	testTreeHas_IsRoot(t)
	testTreeHas_IsChild(t)
	testTreeHas_IsGrandchild(t)
	testTreeHas_PriorityNotInTree(t)
	testTreeHas_DataNotInTree(t)
}

func testMaxInt_DiffVals(t *testing.T) {
	lower := -1
	higher := 0
	maxVal := maxInt(lower, higher)
	if maxVal != higher {
		t.Errorf("maxInt(%d, %d) == %d, expected %d", lower, higher, maxVal, higher)
	}
	maxVal = maxInt(higher, lower)
	if maxVal != higher {
		t.Errorf("maxInt(%d, %d) == %d, expected %d", higher, lower, maxVal, higher)
	}
}

func testMaxInt_SameVals(t *testing.T) {
	val := 3
	maxVal := maxInt(val, val)
	if maxVal != val {
		t.Errorf("maxInt(%d,%d) == %d, expected %d", val, val, maxVal, val)
	}
}

func TestMaxInt(t *testing.T) {
	testMaxInt_DiffVals(t)
	testMaxInt_SameVals(t)
}

func createAvlTreeWithHeight(data string, priority int, height int, left *AvlTree, right *AvlTree) *AvlTree {
	tree := NewAvlTree()
	tree.root = createAvlNode(data, priority)
	tree.height = height
	tree.left = left
	tree.right = right
	return tree
}

func createAvlTree_Leaf(data string, priority int) *AvlTree {
	tree := NewAvlTree()
	tree.root = createAvlNode(data, priority)
	tree.height = 0
	return tree
}

func createAvlNode(data string, priority int) *AvlNode {
	return &AvlNode{data, priority}
}
