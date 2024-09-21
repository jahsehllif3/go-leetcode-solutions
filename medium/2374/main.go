package main

import (
	"fmt"
)

func edgeScore(edges []int) int {
	scores := make([]int, len(edges))
	maxNode := 0
	for i, v := range edges {
		scores[v] += i
		if scores[v] > scores[maxNode] {
			maxNode = v
		}
		if scores[v] == scores[maxNode] {
			if v < maxNode {
				maxNode = v
			}
		}
	}
	//fmt.Println(scores)
	return maxNode
}

func main() {
	fmt.Println(edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5}))
	fmt.Println(edgeScore([]int{3, 3, 3, 0}))
}
