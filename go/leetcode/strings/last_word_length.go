package leetcode

import (
	"strings"
)

func LengthOfLastWord(s string) int {

	words := strings.Fields(s)
	lastWord := strings.Join(words[len(words)-1:], "")

	return len([]rune(lastWord))

}
func LengthOfLastWord2(s string) int {

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if length == 0 && s[i] == ' ' {
			continue
		} else {
			if s[i] == ' ' {
				break
			}
			length++
		}
	}
	return length

}
