package traversal

// 二叉树后序遍历递归实现
func postorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}

		if root.Right != nil {
			dfs(root.Right)
		}
		res = append(res, root.Val)
	}
	dfs(root)
	return res
}
