package longest_continuous_increasing_subsequence

/*
https://leetcode.com/problems/longest-continuous-increasing-subsequence/
Given an unsorted array of integers nums,
return the length of the longest continuous increasing subsequence (i.e. subarray).
The subsequence must be strictly increasing.
A continuous increasing subsequence is defined by two indices l and r (l < r)
such that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]]
and for each l <= i < r, nums[i] < nums[i + 1].
*/

func findLengthOfLCIS(nums []int) (result int) {
	switch len(nums) {
	case 0:
	case 1:
		result = 1
	default:
		result = 1
		current := 1
		for index := 1; index < len(nums); index++ {
			if nums[index] > nums[index-1] {
				current++
			} else {
				if current > result {
					result = current
				}
				current = 1
			}
		}
		if current > result {
			result = current
		}

	}
	return
}
