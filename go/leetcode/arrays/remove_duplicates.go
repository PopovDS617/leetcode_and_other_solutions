package leetcode

func RemoveDuplicates(nums []int) int {
	dict := make(map[int]int)
	res := []int{}

	for _, v := range nums {

		if dict[v] == 0 {
			dict[v]++
			res = append(res, v)
		} else {
			continue
		}

	}

	return len(res)
}

func RemoveDuplicates2(nums []int) int {
	pointer := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[pointer] = nums[i]
			pointer++
		}
	}

	return pointer
}
