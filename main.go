package main

import (
	"fmt"
)

func knapsack(values, weights []int, capacity int) int {
	n := len(values)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weights[i-1]])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, capacity int

	fmt.Scan(&n)

	values := make([]int, n)
	weights := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&values[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	fmt.Scan(&capacity)

	result := knapsack(values, weights, capacity)
	fmt.Printf("%d \n", result)
}
