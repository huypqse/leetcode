package array

func findDuplicates(nums []int) []int {
	result := make(map[int]int)
	var arr []int
	for i := 0; i < len(nums); i++ {
		result[nums[i]]++
	}
	for i := 1; i < len(nums); i++ {
		if result[i] == 2 {
			arr = append(arr, i)
		}
	}
	if result[len(nums)] == 2 {
		arr = append(arr, len(nums))
	}
	return arr
}
