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
	node := avlNode{"A", 5, nil, nil}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &nilNode, expectedCompareResult)
}

func testNodeCompare_SameNode(t *testing.T) {
	node := avlNode{"A", 5, nil, nil}
	expectedCompareResult := 0
	verifyCompareVal(t, &node, &node, expectedCompareResult)
}

func testNodeCompare_OtherIsEquivalent(t *testing.T) {
	node := avlNode{"A", 5, nil, nil}
	equivNode := avlNode{"A", 5, nil, nil}
	expectedCompareResult := 0
	verifyCompareVal(t, &node, &equivNode, expectedCompareResult)
}

func testNodeCompare_OtherHasLowerPriority(t *testing.T) {
	node := avlNode{"A", 5, nil, nil}
	lowerPriorityNode := avlNode{"B", 3, nil, nil}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &lowerPriorityNode, expectedCompareResult)
}

func testNodeCompare_OtherHasHigherPriority(t *testing.T) {
	node := avlNode{"A", 5, nil, nil}
	higherPriorityNode := avlNode{"B", 8, nil, nil}
	expectedCompareResult := -1
	verifyCompareVal(t, &node, &higherPriorityNode, expectedCompareResult)
}

func testNodeCompare_OtherHasSamePriorityLowerData(t *testing.T) {
	node := avlNode{"ABC", 5, nil, nil}
	other := avlNode{"AAA", 5, nil, nil}
	expectedCompareResult := 1
	verifyCompareVal(t, &node, &other, expectedCompareResult)
}

func testNodeCompare_OtherHasSamePriorityHigherData(t *testing.T) {
	node := avlNode{"AAA", 5, nil, nil}
	other := avlNode{"ABC", 5, nil, nil}
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
