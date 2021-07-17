/*
lc 92:https://leetcode-cn.com/problems/reverse-linked-list-ii/
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// left=1时，表示反转链表的第n个节点
	if left == 1 {
		return reverseN(head, right)
	}
	pre, cur := head, head
	for left != 1 {
		pre = cur
		cur = cur.Next
		left--
		right--
	}
	pre.Next = reverseN(cur, right)
	return head
}

func reverseN(node *ListNode, n int) *ListNode {
	if node == nil {
		return nil
	}
	var pre, next *ListNode
	cur := node
	for i := 0; i < n && cur != nil; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	node.Next = cur
	return pre
}
