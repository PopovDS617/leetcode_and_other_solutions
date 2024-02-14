package leetcode

import "sort"

func LastStoneWeight(stones []int) int {

	switch len(stones) {
	case 0:
		return 0
	case 1:
		return stones[0]
	}

	for i := 0; i < len(stones); i++ {

		sort.Slice(stones, func(a, b int) bool {
			return stones[a] > stones[b]
		})

		rest := stones[0] - stones[1]

		if rest > 0 {
			stones[0] = rest
			stones[1] = 0
		}

		if rest == 0 {
			stones[0] = 0
			stones[1] = 0
		}
	}

	for _, v := range stones {

		if v != 0 {
			stones = append(stones, v)
		}
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}
