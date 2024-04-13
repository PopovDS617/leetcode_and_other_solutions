package leetcode

func FindFirstOccerence(str string, need string) int {

	var check string

	if len(need) == len(str) && str == need {
		return 0
	}

	runedNeed := []rune(need)

	for i := range str {

		if i+len(runedNeed) > len(str) {
			break
		}

		if check = str[i : i+len(runedNeed)]; check == need {
			return i
		}

	}

	return -1
}
