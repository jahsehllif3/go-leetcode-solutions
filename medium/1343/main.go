package main

import (
	"fmt"
)

// 定长滑动！！！
// 1. 先计算出第一个定长窗口的和
// 2. 从第一个定长窗口开始，每次移动一个元素，计算新的和，然后和之前的最大和比较
// 2.1 计算新的和=新的和+新的元素-旧的元素
func numOfSubarrays(arr []int, k int, threshold int) int {
	cnt := 0
	arrLen := len(arr)
	sum := 0
	for i := 0; i < k; i++ {
		if i > arrLen {
			break
		}
		sum += arr[i]
	}
	if sum/k >= threshold {
		cnt++
	}
	for j := k; j < arrLen; j++ {
		sum += arr[j] - arr[j-k]
		if sum/k >= threshold {
			cnt++
		}
	}
	return cnt
}

func main() {
	//// 3
	//fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))
	//// 6
	//fmt.Println(numOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5))
	// 1
	fmt.Println(numOfSubarrays([]int{7, 7, 7, 7, 7, 7, 7}, 7, 7))
}
