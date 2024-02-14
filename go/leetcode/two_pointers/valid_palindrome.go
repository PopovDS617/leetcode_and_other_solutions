package leetcode

import (
	"regexp"
	"strings"
)

func IsValidPalindrome(s string) bool {

	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	for i, rightCur := 0, len(s)-1; rightCur > 0; i++ {

		if s[i] != s[rightCur] {
			return false
		}

		rightCur--

	}

	return true

}
