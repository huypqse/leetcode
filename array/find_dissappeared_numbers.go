package array

func findDisappearedNumbers(nums []int) []int {
	result := make(map[int]int)
	var arr []int
	for i := 0; i < len(nums); i++ {
		result[nums[i]]++
	}
	for i := 1; i < len(nums); i++ {
		if result[i] == 0 {
			arr = append(arr, i)
		}
	}
	if result[len(nums)] == 0 {
		arr = append(arr, len(nums))
	}
	return arr
}
