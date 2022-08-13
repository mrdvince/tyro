func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min_price := prices[0]
	profit := 0
	for _, price := range prices[1:] {
		if price < min_price {
			min_price = price
		}
		profit = max(profit, price-min_price)
	}
	return profit

}
