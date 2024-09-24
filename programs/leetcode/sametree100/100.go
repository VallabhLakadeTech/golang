package sametree100

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if (p != nil && q != nil) && (p.Val == q.Val) {

		leftSame := IsSameTree(p.Left, q.Left)
		rightSame := IsSameTree(p.Right, q.Right)
		if leftSame && rightSame {
			return true
		}
		return false

	} else if p == nil && q == nil {
		return true
	}
	return false

}
