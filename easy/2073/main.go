package main

import (
	"fmt"
)

// 解法1
// 因为队首买完票后，会回到队尾
// k花的时间=（轮数-1）*（每轮的人）+最后一轮的位置
// 时间复杂度：O(tickets[k]*len(tickets))
// 空间复杂度：O(2)
func timeRequiredToBuy1(tickets []int, k int) int {
	round := tickets[k]
	cnt := 0
	for i := 0; i < round; i++ {
		for j := 0; j < len(tickets); j++ {
			if tickets[j] > 0 {
				tickets[j]--
				cnt++
				if j == k && tickets[j] == 0 {
					return cnt
				}
			}
		}
	}
	return cnt
}

// 解法2
// 当k买到票时，k前的人至多买到的票数是tickets[k]，k后的人至多买到的票数是tickets[k]-1
// 时间复杂度：O(len(tickets))
// 空间复杂度：O(1)
func timeRequiredToBuy(tickets []int, k int) int {
	cnt := 0
	tk := tickets[k]
	for i, v := range tickets {
		if i <= k {
			cnt += min(v, tk)
		}
		if i > k {
			cnt += min(v, tk-1)
		}
	}
	return cnt
}

func main() {
	// 6
	fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
	// 8
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}
