package avlTree

import (
	"sync"
)

type avlNode struct {
	data     string
	priority int
	left     *avlNode
	right    *avlNode
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
	root *avlNode
	lock sync.RWMutex
}
