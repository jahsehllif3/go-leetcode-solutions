package main

import (
	"fmt"
)

// 不定长滑动窗口
func lengthOfLongestSubstring(s string) int {
	m := [128]bool{}
	left := 0
	r := 0
	for right, c := range s {
		for m[c] {
			m[s[left]] = false
			left++
		}
		m[c] = true
		r = max(r, right-left+1)
	}
	return r
}
func main() {
	fmt.Println(lengthOfLongestSubstring("asdsad"))
	fmt.Println(lengthOfLongestSubstring("bbbc"))
	fmt.Println(lengthOfLongestSubstring("asvxv"))
}
