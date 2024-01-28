package binarytreepaths

import (
	"fmt"
	"strings"
)

/*
https://leetcode.com/problems/binary-tree-paths/description/
Given the root of a binary tree, return all root-to-leaf paths in any order.
A leaf is a node with no children.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(node *TreeNode, path ...string) (out []string) {
	path = append(path, fmt.Sprintf("%v", node.Val))

	if node.Left == nil && node.Right == nil {
		return []string{strings.Join(path, "->")}
	}

	if node.Left != nil {
		out = append(out, binaryTreePaths(node.Left, path...)...)
	}
	if node.Right != nil {
		out = append(out, binaryTreePaths(node.Right, path...)...)
	}

	return
}
