package main

import (
	"fmt"
)

// 依旧是定长滑窗
// 之前都是无脑求出第一个定长窗，然后再滑动
func getAverages(nums []int, k int) []int {
	l := len(nums)
	r := make([]int, l)
	for i, _ := range r {
		r[i] = -1
	}
	// 窗长度
	rL := k*2 + 1
	sum := 0
	for i, v := range nums {
		if i < rL-1 {
			sum += v
		}
		if i >= rL-1 {
			sum += v
			r[i-k] = sum / rL
			sum -= nums[i-rL+1]
		}
	}
	return r
}
func main() {
	// [-1,-1,-1,5,4,4,-1,-1,-1]
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
}
