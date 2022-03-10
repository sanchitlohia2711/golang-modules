package main

import "fmt"

func maxProfit(prices []int) int {
	lenPrices := len(prices)

	buy := -1
	sell := -1

	maxProphit := 0

	for i := 0; i < lenPrices; {
		for i+1 < lenPrices && prices[i] > prices[i+1] {
			i++
		}

		if i == lenPrices-1 {
			return maxProphit
		}

		buy = i

		i++

		for i+1 < lenPrices && prices[i] < prices[i+1] {
			i++
		}

		sell = i

		if (prices[sell] - prices[buy]) > maxProphit {
			maxProphit = prices[sell] - prices[buy]
		}
		i++
	}

	return maxProphit
}

func main() {
	output := maxProfit([]int{4, 2, 3, 8, 1})
	fmt.Println(output)
}
