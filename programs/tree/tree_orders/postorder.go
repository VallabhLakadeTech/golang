package main

import "fmt"

func postorder(tree *TreeNode[int]) {

	if tree != nil {
		inorder(tree.leftNode)
		inorder(tree.rightNode)
		fmt.Println(tree.data)

	}

}
