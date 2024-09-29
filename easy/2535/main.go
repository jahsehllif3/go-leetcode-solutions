package main

import (
	"fmt"
	"strconv"
)

func differenceOfSum(nums []int) int {
	r := 0
	eSum := 0
	dSum := 0
	for _, v := range nums {
		eSum += v
		s := strconv.Itoa(v)
		for _, c := range s {
			dSum += int(c - '0')
		}
	}
	r = eSum - dSum
	if r < 0 {
		r = -r
	}
	return r
}

func main() {
	// 9
	fmt.Println(differenceOfSum([]int{1, 15, 6, 3}))
	// 0
	fmt.Println(differenceOfSum([]int{1, 2, 3, 4}))
}
