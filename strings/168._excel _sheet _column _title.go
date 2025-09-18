package strings

func convertToTitle(num int) string {
	arr := make(map[int]byte, 26)
	var result string
	for i := 0; i < 26; i++ {
		arr[i] = byte('A' + i)
	}
	for num > 0 {
		remain := (num - 1) % 26
		result = string(arr[remain]) + result
		num = (num - 1) / 26

	}
	return result
}
