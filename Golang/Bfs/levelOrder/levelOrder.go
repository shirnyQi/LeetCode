package levelOrder

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 102. 二叉树的层序遍历
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var crtLevel []int
		lenth := queue.Len()
		for i := 0; i < lenth; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			crtLevel = append(crtLevel, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		res = append(res, crtLevel)

	}
	return res

}
