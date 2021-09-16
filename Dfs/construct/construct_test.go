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
