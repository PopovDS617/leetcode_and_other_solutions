package leetcode

func TwoSum(nums []int, target int) []int {
	numMap := map[int]int{}

	for i, num := range nums {
		complement := target - num

		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{}
}

// old solution
func TwoNumberSum(ints []int, target int) []int {

	var result []int

outer:
	for i := 0; i < len(ints); i++ {

		for j := 0; j < len(ints); j++ {

			if ints[i]+ints[j] == target && i != j {
				result = append(result, ints[i], ints[j])
				break outer
			}

		}

	}

	return result
}
