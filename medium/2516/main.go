package main

import (
	"fmt"
)

// 正难则反
// 题目意思取走 k 个 a, b, c 字母，求最少取的次数
// 可以转换为求最长的子串，使得子串中 a, b, c 字母的个数都大于等于 k
// 条件继续转换为 取走的字母的个数都小于 k
func takeCharacters(s string, k int) int {
	cnt := [3]int{}
	for _, c := range s {
		cnt[c-'a']++ // 一开始，把所有字母都取走
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1 // 字母个数不足 k
	}

	r, left := 0, 0
	for right, c := range s {
		cnt[c-'a']--
		for cnt[c-'a'] < k {
			cnt[s[left]-'a']++
			left++
		}
		r = max(r, right-left+1)
	}
	return len(s) - r
}

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2)) // 8
	fmt.Println(takeCharacters("a", 1))            // -1
}
