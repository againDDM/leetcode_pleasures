package leafsimilartrees

/*
https://leetcode.com/problems/leaf-similar-trees/description/
Consider all the leaves of a binary tree, from left to right order,
the values of those leaves form a leaf value sequence.
For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
Two binary trees are considered leaf-similar if their leaf value sequence is the same.
Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dumpToSlice(root *TreeNode, data *[]int) {
	if root.Left == nil && root.Right == nil {
		*data = append(*data, root.Val)
		return
	}

	if root.Left != nil {
		dumpToSlice(root.Left, data)
	}
	if root.Right != nil {
		dumpToSlice(root.Right, data)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	dumpedSlice1 := make([]int, 0, 16)
	dumpedSlice2 := make([]int, 0, 16)
	dumpToSlice(root1, &dumpedSlice1)
	dumpToSlice(root2, &dumpedSlice2)

	if len(dumpedSlice1) != len(dumpedSlice2) {
		return false
	}
	for i, v1 := range dumpedSlice1 {
		if v1 != dumpedSlice2[i] {
			return false
		}
	}
	return true
}
