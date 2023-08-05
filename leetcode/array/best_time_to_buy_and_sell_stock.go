package array

// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
//
// O(n) time | O(1) space
func maxProfit(prices []int) int {
	maxProfit, minPrice := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
