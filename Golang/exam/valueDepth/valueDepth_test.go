package test2

import (
	"fmt"
	"testing"
)

type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

var (
	node1 = &BinaryTreeNode{val: 5, left: node2, right: node3}
	node2 = &BinaryTreeNode{val: 3, left: node4, right: node5}
	node3 = &BinaryTreeNode{val: 7, left: nil, right: nil}
	node4 = &BinaryTreeNode{val: 1, left: nil, right: nil}
	node5 = &BinaryTreeNode{val: 5, left: nil, right: nil}
)

var (
	node6  = &BinaryTreeNode{val: 3, left: node7, right: node8}
	node7  = &BinaryTreeNode{val: 4, left: node9, right: nil}
	node8  = &BinaryTreeNode{val: 11, left: nil, right: node10}
	node9  = &BinaryTreeNode{val: 3, left: node11, right: nil}
	node10 = &BinaryTreeNode{val: 8, left: node12, right: nil}
	node11 = &BinaryTreeNode{val: 7, left: nil, right: nil}
	node12 = &BinaryTreeNode{val: 5, left: nil, right: nil}
)

func valueDepth(target []int, root *BinaryTreeNode) []int {
	var maxDepth = make(map[int]int, 0)
	var queue []*BinaryTreeNode
	queue = append(queue, root)
	depth, maxNode, minNode := 0, 0, 0
	for len(queue) > 0 {
		len := len(queue)
		depth++
		for i := 0; i < len; i++ {
			node := queue[i]
			maxDepth[node.val] = depth
			maxNode = max(maxNode, node.val)
			minNode = min(minNode, node.val)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		queue = queue[len:]
	}
	for i := 0; i < len(target); i++ {
		if target[i] > maxNode || target[i] < minNode {
			target[i] = -1
			continue
		}
		target[i] = getMaxDepth(maxDepth, target[i])
	}
	return target
}

func getMaxDepth(m map[int]int, t int) int {
	var res = -1
	for k, v := range m {
		if k > t && v > res {
			res = v
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestName(t *testing.T) {
	fmt.Println(valueDepth([]int{2, 7}, node1))
	fmt.Println(valueDepth([]int{6, 2, 9, 7, 9}, node6))
}
