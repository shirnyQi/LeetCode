package traversal

// 二叉树前序遍历递归实现
func preorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}
