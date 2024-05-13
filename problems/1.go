package problems

import (
	"encoding/json"
	"fmt"
	"os"
)

func BestPath() int{
	data, err := os.ReadFile("./problems/hard.json")
    if err != nil {
        fmt.Println("error:",err)
    }
	var levels [][]int
	if err := json.Unmarshal(data, &levels); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	// var levels = [][]int{
    //     {59},
    //     {73, 41},
    //     {52, 40, 53},
    //     {26, 53, 6, 34},
    // }
	// fmt.Println(levels)

	dp := make([][]int, len(levels))
	for i := range dp {
		dp[i] = make([]int, len(levels[i]))
	}

	dp[0][0] = levels[0][0]

	for i := 1; i < len(levels); i++ {
		for j := 0; j < len(levels[i]); j++ {
			if j == 0 {
				// first
				dp[i][j] = dp[i-1][j] + levels[i][j]
			} else if j == len(levels[i])-1 {
				// last
				dp[i][j] = dp[i-1][j-1] + levels[i][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]) + levels[i][j]
			}
		}
	}

	maxSum := dp[len(levels)-1][0]
	for j := 1; j < len(levels[len(levels)-1]); j++ {
		maxSum = max(maxSum, dp[len(levels)-1][j])
	}
	fmt.Println(maxSum)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}