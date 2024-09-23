package main

import (
	"fmt"
)

// 两地分数之和 - 两地之间的距离
// values[i] + values[j] + i - j = values[i] + i + values[j] - j
func maxScoreSightseeingPair(values []int) int {
	startScore := 0
	score := 0
	for j, v := range values {
		score = max(score, startScore+v-j)
		startScore = max(startScore, v+j)
	}
	return score
}

func main() {
	// fmt.Println(maxScoreSightseeingPair([]int{1, 2}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 3, 5}))
	// fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
}
