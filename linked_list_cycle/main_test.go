package linked_list_cycle

import (
	"testing"
)

func TestHasCycleNil(t *testing.T) {
	if hasCycle(nil) {
		t.Error("empty linked list can not contain cycle")
	}
}

func TestHasCycleSingleNonCycled(t *testing.T) {
	if hasCycle(&ListNode{}) {
		t.Error("it is non-cycled linked list")
	}
}

func TestHasCycleSingleCycled(t *testing.T) {
	list := ListNode{}
	list.Next = &list
	if !hasCycle(&list) {
		t.Error("it is cycled linked list")
	}
}

func TestHasCycleDoubleNonCycled(t *testing.T) {
	if hasCycle(&ListNode{Next: &ListNode{}}) {
		t.Error("it is non-cycled linked list")
	}
}

func TestHasCycleDoubleCycledToFirst(t *testing.T) {
	first := ListNode{}
	second := ListNode{Next: &first}
	first.Next = &second
	if !hasCycle(&first) {
		t.Error("it is cycled linked list")
	}
}

func TestHasCycleDoubleCycledToSecond(t *testing.T) {
	second := ListNode{}
	first := ListNode{Next: &second}
	second.Next = first.Next
	if !hasCycle(&first) {
		t.Error("it is cycled linked list")
	}
}
