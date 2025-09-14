import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var result strings.Builder
	for i := 0; i < len(arr); i++ {
		word := arr[i]
		for j := len(word) - 1; j >= 0; j-- {
			result.WriteString(string(word[j]))
		}
		if i < len(arr)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}