package merge_two_sorted_lists

import "testing"

func TestMergeTwoLists1(t *testing.T) {
	list1 := listNodeFromSlice([]int{1, 2, 4})
	list2 := listNodeFromSlice([]int{1, 3, 4})
	expectedResultList := listNodeFromSlice([]int{1, 1, 2, 3, 4, 4})
	obtainedResult := mergeTwoLists(list1, list2)
	if !listsEquals(expectedResultList, obtainedResult) {
		t.Errorf("expect %v got %v", listToSlice(expectedResultList), listToSlice(obtainedResult))
	}
}

func TestMergeTwoLists2(t *testing.T) {
	list1 := listNodeFromSlice([]int{})
	list2 := listNodeFromSlice([]int{})
	expectedResultList := listNodeFromSlice([]int{})
	obtainedResult := mergeTwoLists(list1, list2)
	if !listsEquals(expectedResultList, obtainedResult) {
		t.Errorf("expect %v got %v", listToSlice(expectedResultList), listToSlice(obtainedResult))
	}
}

func TestMergeTwoLists3(t *testing.T) {
	list1 := listNodeFromSlice([]int{})
	list2 := listNodeFromSlice([]int{0})
	expectedResultList := listNodeFromSlice([]int{0})
	obtainedResult := mergeTwoLists(list1, list2)
	if !listsEquals(expectedResultList, obtainedResult) {
		t.Errorf("expect %v got %v", listToSlice(expectedResultList), listToSlice(obtainedResult))
	}
}

func TestMergeTwoLists4(t *testing.T) {
	list1 := listNodeFromSlice([]int{2})
	list2 := listNodeFromSlice([]int{1})
	expectedResultList := listNodeFromSlice([]int{1, 2})
	obtainedResult := mergeTwoLists(list1, list2)
	if !listsEquals(expectedResultList, obtainedResult) {
		t.Errorf("expect %v got %v", listToSlice(expectedResultList), listToSlice(obtainedResult))
	}
}
