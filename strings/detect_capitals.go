package strings

import (
	"unicode"
)

func detectCapitalUse(word string) bool {
	n := len(word)
	count := 0
	for _, ch := range word {
		if unicode.IsUpper(ch) {
			count++
		}
	}
	return count == n || count == 0 || (count == 1 && unicode.IsUpper(rune(word[0])))

}
