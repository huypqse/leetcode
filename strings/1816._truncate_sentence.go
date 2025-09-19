package strings

import "strings"

func truncateSentence(s string, k int) string {
	var result string
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr); i++ {
		if i < k-1 {
			result += string(arr[i]) + " "
		}
		if i == k-1 {
			result += string(arr[i])
			break
		}
	}
	return result
}
