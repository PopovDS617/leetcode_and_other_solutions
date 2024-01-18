package main

// var ints = []int{3, 5, -4, 8, 11, 1, -1, 6}
// var target = 10

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
