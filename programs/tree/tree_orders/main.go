package main

type TreeNode[T any] struct {
	leftNode, rightNode *TreeNode[T]
	data                T
}

func main() {

	tree := TreeNode[int]{
		leftNode:  nil,
		rightNode: nil,
		data:      0,
	}
	inorder(tree)
}
