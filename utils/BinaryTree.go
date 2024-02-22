package utils

import "fmt"

type BinaryTreeNode struct {
	Data int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewNode(value int) *BinaryTreeNode {
	node := &BinaryTreeNode{
		Data: value,
		Left: nil,
		Right: nil,
	}
	return node;
}

func PrintInorder(node *BinaryTreeNode) {
	if node == nil {
		return;
	}
	PrintInorder(node.Left)
	fmt.Printf("%d  ", node.Data);
	PrintInorder(node.Right);
}

