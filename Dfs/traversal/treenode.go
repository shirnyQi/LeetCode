package traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	node1 = &TreeNode{1, nil, nil}
	node5 = &TreeNode{5, nil, nil}
	node0 = &TreeNode{0, nil, nil}
	node9 = &TreeNode{9, node5, node1}
	node4 = &TreeNode{4, node9, node0}
)
