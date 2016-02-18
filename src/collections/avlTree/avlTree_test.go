package avlTree

import (
	// "fmt"
	"runtime/debug"
	"testing"
)

func verifyCompareVal(t *testing.T, node *avlNode, other *avlNode, expectedVal int) {
	compareVal := node.compare(other)
	if compareVal != expectedVal {
		t.Errorf("%v.compare(%v) == %d, expected %d", *node, *other, compareVal, expectedVal)
		debug.PrintStack()
	}
}

func testNodeCompare_OtherIsNil(t *testing.T) {
	var nilNode avlNode
	node := avlNode{"A", 5}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &nilNode, expectedCompareResult)
}

func testNodeCompare_SameNode(t *testing.T) {
	node := avlNode{"A", 5}
	expectedCompareResult := 0
	verifyCompareVal(t, &node, &node, expectedCompareResult)
}

func testNodeCompare_OtherIsEquivalent(t *testing.T) {
	node := avlNode{"A", 5}
	equivNode := avlNode{"A", 5}
	expectedCompareResult := 0
	verifyCompareVal(t, &node, &equivNode, expectedCompareResult)
}

func testNodeCompare_OtherHasLowerPriority(t *testing.T) {
	node := avlNode{"A", 5}
	lowerPriorityNode := avlNode{"B", 3}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &lowerPriorityNode, expectedCompareResult)
}

func testNodeCompare_OtherHasHigherPriority(t *testing.T) {
	node := avlNode{"A", 5}
	higherPriorityNode := avlNode{"B", 8}
	expectedCompareResult := -1
	verifyCompareVal(t, &node, &higherPriorityNode, expectedCompareResult)
}

func testNodeCompare_OtherHasSamePriorityLowerData(t *testing.T) {
	node := avlNode{"ABC", 5}
	other := avlNode{"AAA", 5}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &other, expectedCompareResult)
}

func testNodeCompare_OtherHasSamePriorityHigherData(t *testing.T) {
	node := avlNode{"AAA", 5}
	other := avlNode{"ABC", 5}
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

	emptyTree := newAvlTree()
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
	rootNode := createAvlNode("higher", 3)
	tree := createAvlTree(rootNode, left, nil)

	expectedHeight := 1
	verifyTreeCalcHeightFromChildrenVal(t, tree, expectedHeight)
}

func testTreeCalcHeightFromChildren_Grandparent(t *testing.T) {
	grandchild := createAvlTree_Leaf("low", 2)
	left := createAvlTreeWithHeight("lowest", 1, 1, nil, grandchild)
	rootNode := createAvlNode("high", 3)
	tree := createAvlTree(rootNode, left, nil)

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
		t.Errorf("getHeight() == %d, expected %d", tree, height, expectedHeight)
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
	leaf.Insert(leafNode)
	expectedHeight := 0
	verifyGetHeightVal(t, &leaf, expectedHeight)
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

	emptyTree := newAvlTree()
	verifyUpdateHeight(t, emptyTree, expectedNewHeight)
}

func testTreeUpdateHeight_Leaf(t *testing.T) {
	leaf := createAvlTree_Leaf("a", 1)

	expectedNewHeight := 0
	verifyUpdateHeight(t, leaf, expectedNewHeight)
}

func testTreeUpdateHeight_Parent(t *testing.T) {
	leaf := createAvlTree_Leaf("a", 1)

	parentNode := createAvlNode("b", 5)
	parent := createAvlTree(parentNode, leaf, nil)

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
		verifyTreeNodesEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		rotateLeftToRoot(&tree)
		verifyTreeNodesEqual(t, tree, prevTree)
	}
}

func verifyTreeRotateRight_Empty(t *testing.T, tree *AvlTree) {
	if tree == nil {
		rotateRightToRoot(&tree)
		verifyTreeNodesEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		rotateRightToRoot(&tree)
		verifyTreeNodesEqual(t, tree, prevTree)
	}
}

func verifyTreeNodesEqual(t *testing.T, tree *AvlTree, expected *AvlTree) {
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
	verifyTreeNodesEqual(t, tree.left, expectedLeft)
	verifyTreeNodesEqual(t, tree.right, expectedRight)
}

func testTreeRotateLeft_EmptyTree(t *testing.T) {
	verifyTreeRotateLeft_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeRotateLeft_Empty(t, &nilTree)

	emptyTree := newAvlTree()
	verifyTreeRotateLeft_Empty(t, emptyTree)
}

func testTreeRotateRight_EmptyTree(t *testing.T) {
	verifyTreeRotateRight_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeRotateRight_Empty(t, &nilTree)

	emptyTree := newAvlTree()
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
		verifyTreeNodesEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		doubleRotateLeftToRoot(&tree)
		verifyTreeNodesEqual(t, tree, prevTree)
	}
}

func verifyTreeDoubleRotateRight_Empty(t *testing.T, tree *AvlTree) {
	if tree == nil {
		doubleRotateRightToRoot(&tree)
		verifyTreeNodesEqual(t, tree, nil)
	} else {
		prevTree := &*tree
		doubleRotateRightToRoot(&tree)
		verifyTreeNodesEqual(t, tree, prevTree)
	}
}

func testTreeDoubleRotateLeft_EmptyTree(t *testing.T) {
	verifyTreeDoubleRotateLeft_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeDoubleRotateLeft_Empty(t, &nilTree)

	emptyTree := newAvlTree()
	verifyTreeDoubleRotateLeft_Empty(t, emptyTree)
}

func testTreeDoubleRotateRight_EmptyTree(t *testing.T) {
	verifyTreeDoubleRotateRight_Empty(t, nil)

	var nilTree AvlTree
	verifyTreeDoubleRotateRight_Empty(t, &nilTree)

	emptyTree := newAvlTree()
	verifyTreeDoubleRotateRight_Empty(t, emptyTree)
}

func testTreeDoubleRotateLeft_ParentWithNoLeft(t *testing.T) {
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, nil, right)
	prevTree := &*tree
	doubleRotateLeftToRoot(&tree)
	verifyTreeNodesEqual(t, tree, prevTree)
}

func testTreeDoubleRotateRight_ParentWithNoRight(t *testing.T) {
	left := createAvlTree_Leaf("left", 2)
	tree := createAvlTreeWithHeight("parent", 5, 1, left, nil)
	prevTree := &*tree
	doubleRotateRightToRoot(&tree)
	verifyTreeNodesEqual(t, tree, prevTree)
}

func testTreeDoubleRotateLeft_LeftWithNoRight(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevTree := &*tree
	doubleRotateLeftToRoot(&tree)
	verifyTreeNodesEqual(t, tree, prevTree)
}

func testTreeDoubleRotateRight_RightWithNoLeft(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	tree := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevTree := &*tree
	doubleRotateRightToRoot(&tree)
	verifyTreeNodesEqual(t, tree, prevTree)
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
	verifyTreeNodesEqual(t, tree, prevTree)
}

func testTreeBalance_Nil(t *testing.T) {
	var tree *AvlTree = nil
	balance(&tree)
	verifyTreeNodesEqual(t, tree, nil)
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

	emptyTree := newAvlTree()
	verifyTreeBalanceHasNoEffect(t, emptyTree)
}

func testTreeBalance_BalancesIn1LftRtt(t *testing.T) {
	leftL := createAvlTree_Leaf("LL", -3)
	left := createAvlTreeWithHeight("L", -1, 1, leftL, nil)
	tree := createAvlTreeWithHeight("root", 2, 2, left, nil)
	prevTree := &*tree

	balance(&tree)
	verifyTreeNodesEqual(t, tree, left)
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
	verifyTreeNodesEqual(t, tree, right)
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

func testMax_DiffVals(t *testing.T) {
	lower := -1
	higher := 0
	maxVal := max(lower, higher)
	if maxVal != higher {
		t.Errorf("max(%d, %d) == %d, expected %d", lower, higher, maxVal, higher)
	}
	maxVal = max(higher, lower)
	if maxVal != higher {
		t.Errorf("max(%d, %d) == %d, expected %d", higher, lower, maxVal, higher)
	}
}

func testMax_SameVals(t *testing.T) {
	val := 3
	maxVal := max(val, val)
	if maxVal != val {
		t.Errorf("max(%d,%d) == %d, expected %d", val, val, maxVal, val)
	}
}

func TestMax(t *testing.T) {
	testMax_DiffVals(t)
	testMax_SameVals(t)
}

func createAvlTree(rootNode *avlNode, left *AvlTree, right *AvlTree) *AvlTree {
	tree := newAvlTree()
	tree.root = rootNode
	tree.left = left
	tree.right = right
	return tree
}

func createAvlTreeWithHeight(data string, priority int, height int, left *AvlTree, right *AvlTree) *AvlTree {
	rootNode := createAvlNode(data, priority)
	tree := createAvlTree(rootNode, left, right)
	tree.height = height
	return tree
}

func createAvlTree_Leaf(data string, priority int) *AvlTree {
	node := createAvlNode(data, priority)
	return createAvlTree(node, nil, nil)
}

func createAvlNode(data string, priority int) *avlNode {
	return &avlNode{data, priority}
}
