package main

import "fmt"

func preorder(tree *TreeNode[int]) {

	if tree != nil {
		fmt.Println(tree.data)
		inorder(tree.leftNode)
		inorder(tree.rightNode)
	}

}
