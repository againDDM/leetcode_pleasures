package longest_continuous_increasing_subsequence

import "testing"

func TestFindLengthOfLCIS1(t *testing.T) {
	if result := findLengthOfLCIS([]int{1, 3, 5, 4, 7}); result != 3 {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, 3)
	}
}

func TestFindLengthOfLCIS2(t *testing.T) {
	if result := findLengthOfLCIS([]int{2, 2, 2, 2, 2}); result != 1 {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, 1)
	}
}
