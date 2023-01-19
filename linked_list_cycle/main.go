package linked_list_cycle

/*
https://leetcode.com/problems/linked-list-cycle/
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list
that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle0(head *ListNode) bool {
	if head == nil {
		return false
	}
	pointersDict := make(map[*ListNode]struct{}, 2048)
	for current := head; current.Next != nil; current = current.Next {
		if _, exists := pointersDict[current]; exists {
			return true
		}
		pointersDict[current] = struct{}{}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycleJoke(head *ListNode) bool {
	if head == nil {
		return false
	}
	for current, count := head, 0; current.Next != nil; current, count = current.Next, count+1 {
		if count > 10000 {
			return true
		}
	}
	return false
}
