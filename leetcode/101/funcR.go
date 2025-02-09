package main

func isSymmetricR(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return solve(root.Left, root.Right)
}

func solve(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	} else {
		return l.Val == r.Val && solve(l.Left, r.Right) && solve(l.Right, r.Left)
	}
}
