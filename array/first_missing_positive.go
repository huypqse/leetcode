package array

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	result := make(map[int]int)
	var arr int
	max := 0
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		result[nums[i]]++
	}
	if len(nums) > max {
		max = len(nums)
		if result[max] == 0 {
			arr = max
		}
	}
	for i := 1; i < max; i++ {
		if result[i] == 0 {
			arr = i
			break
		}
	}
	if arr == 0 {
		arr = len(nums) + 1
	}
	return arr
}
