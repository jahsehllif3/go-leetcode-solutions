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

func main() {
	fmt.Println(maxVowels("abciiidef", 3))    // 3
	fmt.Println(maxVowels("aeiou", 2))        // 2
	fmt.Println(maxVowels("weallloveyou", 7)) // 4
}
