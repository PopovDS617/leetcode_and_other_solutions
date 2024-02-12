package leetcode

import (
	"bytes"
	"sort"
)

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	bytedS := []byte(s)
	bytedT := []byte(t)

	sort.Slice(bytedT, func(i, j int) bool {
		return bytedT[i] < bytedT[j]
	})
	sort.Slice(bytedS, func(i, j int) bool {
		return bytedS[i] < bytedS[j]
	})

	return bytes.Equal(bytedS, bytedT)
}
