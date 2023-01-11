package merge_two_sorted_lists

/*
https://leetcode.com/problems/merge-two-sorted-lists/
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list.
The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/

// ListNode is a definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func listNodeFromSlice(srcSlice []int) (list *ListNode) {
	for i := len(srcSlice) - 1; i >= 0; i-- {
		list = &ListNode{Val: srcSlice[i], Next: list}
	}
	return list
}

func listToSlice(ln *ListNode) (result []int) {
	for current := ln; current != nil; {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func listsEquals(l *ListNode, r *ListNode) bool {
	switch true {
	case l == nil && r == nil:
		return true
	case l != nil && r != nil && l.Val == r.Val:
		return listsEquals(l.Next, r.Next)
	default:
		return false
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	switch true {
	case list1 == nil:
		return list2
	case list2 == nil:
		return list1
	}
	resultHead := &ListNode{}
	if list1.Val < list2.Val {
		resultHead.Val = list1.Val
		list1 = list1.Next
	} else {
		resultHead.Val = list2.Val
		list2 = list2.Next
	}
	currentNodeResultPointer := resultHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNodeResultPointer.Next = &ListNode{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else {
			currentNodeResultPointer.Next = &ListNode{Val: list2.Val, Next: nil}
			list2 = list2.Next
		}
		currentNodeResultPointer = currentNodeResultPointer.Next
	}
	if list1 != nil {
		currentNodeResultPointer.Next = list1
	} else {
		currentNodeResultPointer.Next = list2
	}
	return resultHead
}
