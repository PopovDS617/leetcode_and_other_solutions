package leetcode

func FindUniq(arr []float32) float32 {

	res := make(map[float32]int)

	for _, v := range arr {
		res[v]++
	}

	for k, v := range res {
		if v == 1 {
			return k
		}
	}
	return arr[0]
}
