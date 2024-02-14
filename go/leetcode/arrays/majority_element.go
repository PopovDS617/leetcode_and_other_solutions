package leetcode

func MajorityElement(nums []int) int {

	store := map[int]int{}

	majVal := 0
	majCount := 0

	for i := 0; i < len(nums); i++ {

		store[nums[i]]++

		if x := store[nums[i]]; x > majCount {
			majVal = nums[i]
			majCount = x
		}

	}
	return majVal
}
