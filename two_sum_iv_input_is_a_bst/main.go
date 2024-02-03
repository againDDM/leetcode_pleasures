package twosumivinputisabst

/*
https://leetcode.com/problems/two-sum-iv-input-is-a-bst
Given the root of a binary search tree and an integer k, return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.
Input: root = [5,3,6,2,4,null,7], k = 9
Output: true
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(current *TreeNode, cache map[int]struct{}, k int) bool {
	if current == nil {
		return false
	}
	if _, cached := cache[k-current.Val]; cached {
		return true
	}
	cache[current.Val] = struct{}{}
	return dfs(current.Left, cache, k) || dfs(current.Right, cache, k)
}

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, make(map[int]struct{}, 64), k)
}
