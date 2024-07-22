package main

import "fmt"

func inorder(tree *TreeNode[int]) {

	if tree != nil {
		inorder(tree.leftNode)
		fmt.Println(tree.data)
		inorder(tree.rightNode)
	}

}
