package avlTree

import (
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
	prevTree := tree
	tree.rotateLeft()
	verifyTreeNodesEqual(t, tree, prevTree)
}

func verifyTreeRotateRight_Empty(t *testing.T, tree *AvlTree) {
	prevTree := tree
	tree.rotateRight()
	verifyTreeNodesEqual(t, tree, prevTree)
}

func verifyTreeNodesEqual(t *testing.T, tree *AvlTree, expected *AvlTree) {
	if tree != expected {
		if tree == nil {
			t.Errorf("tree == nil, expected %v\nexpected.root == &%v", expected, expected.root)
			debug.PrintStack()
		} else {
			t.Errorf("tree == %v, expected %v\ntree.root == &%v", tree, expected, tree.root)
			debug.PrintStack()
		}
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
	leaf.rotateLeft()
	if *leaf != prevLeafVal {
		t.Errorf("tree == &%v, expected &%v", *leaf, prevLeafVal)
	}
}

func testTreeRotateRight_Leaf(t *testing.T) {
	leaf := createAvlTree_Leaf("a", 1)
	prevLeafVal := *leaf
	leaf.rotateRight()
	if *leaf != prevLeafVal {
		t.Errorf("tree == &%v, expected &%v", *leaf, prevLeafVal)
	}
}

func testTreeRotateLeft_ParentWithNoLeft(t *testing.T) {
	right := createAvlTree_Leaf("right", 2)
	parentNode := createAvlNode("parent", 1)
	parent := createAvlTree(parentNode, nil, right)
	prevParentVal := *parent
	parent.rotateLeft()
	if *parent != prevParentVal {
		t.Errorf("tree == &%v, expected &%v", *parent, prevParentVal)
	}
}

func testTreeRotateRight_ParentWithNoRight(t *testing.T) {
	left := createAvlTree_Leaf("left", 2)
	parent := createAvlTreeWithHeight("parent", 5, 1, left, nil)
	prevParentVal := *parent
	parent.rotateRight()
	if *parent != prevParentVal {
		t.Errorf("tree == &%v, expected &%v", *parent, prevParentVal)
	}
}

func testTreeRotateLeft_ParentWithLAndR(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	parent := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevLR := left.right
	parent.rotateLeft()

	verifyTreeLAndR(t, parent, prevLR, right)
	verifyTreeLAndR(t, left, nil, parent)
	if left.height != 2 {
		t.Errorf("left.height == %d, expected 1", left.height)
	}
	if parent.height != 1 {
		t.Errorf("parent.height == %d, expected 0", parent.height)
	}
}

func testTreeRotateRight_ParentWithLAndR(t *testing.T) {
	left := createAvlTree_Leaf("left", 0)
	right := createAvlTree_Leaf("right", 2)
	parent := createAvlTreeWithHeight("parent", 1, 1, left, right)
	prevRL := right.left
	parent.rotateRight()

	verifyTreeLAndR(t, parent, left, prevRL)
	verifyTreeLAndR(t, right, parent, nil)
	if right.height != 2 {
		t.Errorf("right.height == %d, expected 1", right.height)
	}
	if parent.height != 1 {
		t.Errorf("parent.height == %d, expected 0", parent.height)
	}
}

func testTreeRotateLeft_LongLeftTail(t *testing.T) {
	tL4 := createAvlTree_Leaf("LLLL", -25)
	tL3 := createAvlTreeWithHeight("LLL", -20, 1, tL4, nil)
	tL2 := createAvlTreeWithHeight("LL", -15, 2, tL3, nil)
	tLR := createAvlTree_Leaf("LR", -5)
	tL := createAvlTreeWithHeight("L", -10, 3, tL2, tLR)
	tR := createAvlTree_Leaf("R", 5)
	root := createAvlTreeWithHeight("root", 0, 4, tL, tR)
	root.rotateLeft()

	verifyTreeLAndR(t, root, tLR, tR)
	verifyTreeLAndR(t, tL, tL2, root)
	expectedLeftHeight := 3
	if tL.height != expectedLeftHeight {
		t.Errorf("left.height == %d, expected %d", tL.height, expectedLeftHeight)
	}
	expectedRootHeight := 1
	if root.height != expectedRootHeight {
		t.Errorf("root.height == %d, expected %d", root.height, expectedRootHeight)
	}
}

func testTreeRotateRight_LongRightTail(t *testing.T) {
	tR4 := createAvlTree_Leaf("RRRR", 25)
	tR3 := createAvlTreeWithHeight("RRR", 20, 1, nil, tR4)
	tR2 := createAvlTreeWithHeight("RR", 15, 2, nil, tR3)
	tRL := createAvlTree_Leaf("R", 5)
	tR := createAvlTreeWithHeight("R", 10, 3, tRL, tR2)
	tL := createAvlTree_Leaf("L", -10)
	root := createAvlTreeWithHeight("root", 0, 4, tL, tR)
	root.rotateRight()

	verifyTreeLAndR(t, root, tL, tRL)
	verifyTreeLAndR(t, tR, root, tR2)
	expectedRightHeight := 3
	if tR.height != expectedRightHeight {
		t.Errorf("right.height == %d, expected %d", tR.height, expectedRightHeight)
	}
	expectedRootHeight := 1
	if root.height != expectedRootHeight {
		t.Errorf("root.height == %d, expected %d", root.height, expectedRootHeight)
	}
}

func TestTreeRotateLeft(t *testing.T) {
	testTreeRotateLeft_EmptyTree(t)
	testTreeRotateLeft_Leaf(t)
	testTreeRotateLeft_ParentWithNoLeft(t)
	testTreeRotateLeft_ParentWithLAndR(t)
	testTreeRotateLeft_LongLeftTail(t)
}

func TestTreeRotateRight(t *testing.T) {
	testTreeRotateRight_EmptyTree(t)
	testTreeRotateRight_Leaf(t)
	testTreeRotateRight_ParentWithNoRight(t)
	testTreeRotateRight_ParentWithLAndR(t)
	testTreeRotateRight_LongRightTail(t)
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
