package leetcode

func ContainsDuplicate(nums []int) bool {
	res := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		res[nums[i]]++
		if res[nums[i]] > 1 {
			return true
		}

	}
	return false
}
