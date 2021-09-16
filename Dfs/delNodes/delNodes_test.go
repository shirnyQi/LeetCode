package delNodes

import (
	"fmt"
	"testing"
)

/*
1110. 删点成林
给出二叉树的根节点 root，树上每个节点都有一个不同的值。
如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
返回森林中的每棵树。你可以按任意顺序组织答案。

示例：

输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
输出：[[1,2,null,4],[6],[7]]

link:https://leetcode-cn.com/problems/delete-nodes-and-return-forest/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	node1 = &TreeNode{1, node2, node3}
	node2 = &TreeNode{2, node4, node5}
	node3 = &TreeNode{3, node6, node7}
	node4 = &TreeNode{4, nil, nil}
	node5 = &TreeNode{5, nil, nil}
	node6 = &TreeNode{6, nil, nil}
	node7 = &TreeNode{7, nil, nil}
)

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var res []*TreeNode

	var m = make(map[int]bool, 0)
	for _, v := range to_delete {
		m[v] = true
	}

	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)
		if !m[root.Val] {
			return root
		}
		if root.Left != nil {
			res = append(res, root.Left)
		}
		if root.Right != nil {
			res = append(res, root.Right)
		}
		return nil
	}
	root = dfs(root)
	if root != nil {
		res = append(res, root)
	}
	return res
}

func TestName(t *testing.T) {
	fmt.Println(delNodes(node1, []int{3, 5}))
}
