/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * link:https://leetcode-cn.com/problems/find-nearest-right-node-in-binary-tree/
 */
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		lenth := queue.Len()
		for i := 0; i < lenth; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
	
            if node == u {
                if i == lenth -1 {
                    return nil
                }
                return queue.Remove(queue.Back()).(*TreeNode)
            }
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
	}
	return nil
}
