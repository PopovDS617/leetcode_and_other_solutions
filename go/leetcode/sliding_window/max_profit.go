package slidingwindow

func MaxProfit(prices []int) int {

	left, right := 0, 1

	maxP := 0

	for right < len(prices) {

		if prices[left] < prices[right] {

			profit := prices[right] - prices[left]

			if profit > maxP {
				maxP = profit
			}

		} else {
			left = right
		}
		right++

	}

	return maxP
}
