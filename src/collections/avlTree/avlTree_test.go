package avlTree

import "testing"

func verifyCompareVal(t *testing.T, node *avlNode, other *avlNode, expectedVal int) {
	compareVal := node.compare(other)
	if compareVal != expectedVal {
		t.Errorf("%v.compare(%v) == %d, expected %d", *node, *other, compareVal, expectedVal)
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
	leafNode := createAvlNode("a", 1)
	leaf := createAvlTree(leafNode, nil, nil)
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
	}
}

func testTreeCalcHeightFromChildren_NilTree(t *testing.T) {
	expectedHeight := -1
	verifyTreeCalcHeightFromChildrenVal(t, nil, expectedHeight)

	var nilTree AvlTree
	verifyTreeCalcHeightFromChildrenVal(t, &nilTree, expectedHeight)
}

func testTreeCalcHeightFromChildren_Leaf(t *testing.T) {
	leafNode := createAvlNode("alpha", 3)
	leaf := createAvlTree(leafNode, nil, nil)
	expectedHeight := 0
	verifyTreeCalcHeightFromChildrenVal(t, leaf, expectedHeight)
}

func testTreeCalcHeightFromChildren_Parent(t *testing.T) {
	leftNode := createAvlNode("lower", 1)
	left := createAvlTree(leftNode, nil, nil)
	rootNode := createAvlNode("higher", 3)
	tree := createAvlTree(rootNode, left, nil)

	expectedHeight := 1
	verifyTreeCalcHeightFromChildrenVal(t, tree, expectedHeight)
}

func testTreeCalcHeightFromChildren_Grandparent(t *testing.T) {
	grandchildNode := createAvlNode("low", 2)
	grandchildTree := createAvlTree(grandchildNode, nil, nil)
	childNode := createAvlNode("lowest", 1)
	left := createAvlTreeWithHeight(childNode, 1, nil, grandchildTree)
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
	height := getHeight(tree)
	if height != expectedHeight {
		t.Errorf("getHeight(%v) == %d, expected %d", tree, height, expectedHeight)
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
	updatedHeight := getHeight(tree)
	if updatedHeight != expectedNewHeight {
		t.Errorf("getHeight(&%v) == %d, expected %d", *tree, updatedHeight, expectedNewHeight)
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
	leafNode := createAvlNode("a", 1)
	leaf := createAvlTree(leafNode, nil, nil)

	expectedNewHeight := 0
	verifyUpdateHeight(t, leaf, expectedNewHeight)
}

func testTreeUpdateHeight_Parent(t *testing.T) {
	leafNode := createAvlNode("a", 1)
	leaf := createAvlTree(leafNode, nil, nil)

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

func createAvlTreeWithHeight(rootNode *avlNode, height int, left *AvlTree, right *AvlTree) *AvlTree {
	tree := createAvlTree(rootNode, left, right)
	tree.height = height
	return tree
}

func createAvlNode(data string, priority int) *avlNode {
	return &avlNode{data, priority}
}
