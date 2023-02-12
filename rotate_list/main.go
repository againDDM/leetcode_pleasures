package rotate_list

/*
https://leetcode.com/problems/rotate-list/
Given the head of a linked list, rotate the list to the right by k places.
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	length := 1
	tail := head
	for ; tail.Next != nil; tail = tail.Next {
		length++
	}
	k = length - k%length
	if k == 0 {
		return head
	}

	tail.Next = head
	for ; k != 0; k-- {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil
	return head
}
