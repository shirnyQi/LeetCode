package construct

import (
	"fmt"
	"testing"
)

/*
427. 建立四叉树
<题目太长，点连接>
link:https://leetcode-cn.com/problems/construct-quad-tree/
*/

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

//func construct(grid [][]int) *Node {
//	size := len(grid)
//	var root *Node
//
//	var dfs func(node *Node, begin, end []int) *Node
//	dfs = func(node *Node, begin, end []int) *Node {
//		start := grid[begin[0]][begin[1]]
//		var flag = true
//		for i := begin[0]; i < end[0]; i++ {
//			for j := begin[1]; j < end[1]; j++ {
//				if grid[i][j] != start {
//					flag = false
//					break
//				}
//			}
//		}
//		if flag {
//			node = &Node{Val: toBool(start), IsLeaf: true}
//			return node
//		}
//		node = &Node{Val: true, IsLeaf: false}
//		node.TopLeft = dfs(node.TopLeft, begin, []int{(begin[0] + end[0]) / 2, (begin[1] + end[1]) / 2})
//		node.TopRight = dfs(node.TopRight, []int{(begin[0] + end[0]) / 2, (begin[1] + end[1]) / 2}, []int{(begin[0] + end[0]) / 2, end[1]})
//		node.BottomLeft = dfs(node.BottomLeft, begin, []int{end[0] / 2, end[1] / 2})
//		node.BottomRight = dfs(node.BottomRight, begin, []int{end[0] / 2, end[1] / 2})
//		return node
//	}
//
//	root = dfs(root, []int{0, 0}, []int{size, size})
//	return root
//}

func construct(grid [][]int) *Node {

	var helper func(x1, y1, x2, y2 int) *Node
	helper = func(x1, y1, x2, y2 int) *Node {
		if x1 == x2-1 {
			val := false
			if grid[x1][y1] == 1 {
				val = true
			}
			return &Node{
				Val:    val,
				IsLeaf: true,
			}
		}
		midX, midY := (x1+x2)>>1, (y1+y2)>>1
		tl := helper(x1, y1, midX, midY)
		tr := helper(x1, midY, midX, y2)
		bl := helper(midX, y1, x2, midY)
		br := helper(midX, midY, x2, y2)
		// 都是叶节点且值一样就丢弃(只保留叶节点的值)
		if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf &&
			(tl.Val && tr.Val && bl.Val && br.Val || (!tl.Val && !tr.Val && !bl.Val && !br.Val)) {
			return &Node{
				Val:    tl.Val,
				IsLeaf: true,
			}
		}
		return &Node{
			IsLeaf:      false,
			TopLeft:     tl,
			TopRight:    tr,
			BottomLeft:  bl,
			BottomRight: br,
		}
	}

	return helper(0, 0, len(grid), len(grid[0]))
}

func toBool(a int) bool {
	if a == 0 {
		return false
	}
	return true
}

func TestName(t *testing.T) {
	fmt.Println(construct([][]int{{0, 1}, {1, 0}}))
	fmt.Println(construct([][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0}}))
	fmt.Println(construct([][]int{{1, 1}, {1, 1}}))
	fmt.Println(construct([][]int{{0}}))
	fmt.Println(construct([][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 1}}))
}
