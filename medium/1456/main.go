package main

import (
	"fmt"
)

// 暴力枚举
// 时间复杂度：O(n*k)
// 空间复杂度：O(1)
//func maxVowels(s string, k int) int {
//	vowels := map[byte]bool{
//		'a': true,
//		'e': true,
//		'i': true,
//		'o': true,
//		'u': true,
//	}
//	left, r, maxR := 0, 0, 0
//	for right := left + k - 1; right < len(s); right++ {
//		for i := right - k + 1; i <= right; i++ {
//			if vowels[s[i]] {
//				r++
//			}
//		}
//		maxR = max(maxR, r)
//		r = 0
//	}
//	return maxR
//}

// 定滑动，优化了每次枚举时，对每个子串的遍历
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	r, maxR := 0, 0
	for i, in := range s {
		if vowels[byte(in)] {
			r++
		}
		if i < k-1 {
			continue
		}
		maxR = max(r, maxR)
		out := s[i-k+1]
		if vowels[out] {
			r--
		}
	}
	return maxR
}

// 定长滑窗套路
// 我总结成三步：入-更新-出。

// 入：下标为 i 的元素进入窗口，更新相关统计量。如果 i<k−1 则重复第一步。
// 更新：更新答案。一般是更新最大值/最小值。
// 出：下标为 i−k+1 的元素离开窗口，更新相关统计量。

func main() {
	fmt.Println(maxVowels("abciiidef", 3))    // 3
	fmt.Println(maxVowels("aeiou", 2))        // 2
	fmt.Println(maxVowels("weallloveyou", 7)) // 4
}
