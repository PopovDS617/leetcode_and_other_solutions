package leetcode

func MoveZeroes(nums []int) {
	pointer := 0
	i := 0

	for i = range nums {

		if nums[i] != 0 {
			nums[pointer] = nums[i]
			pointer++
		}
	}

	for i = pointer; i < len(nums); i++ {
		nums[i] = 0
	}
}
