package main

import (
	"fmt"
)

func longestContinuousSubstring(s string) int {
	// 长度为0
	if len(s) == 0 {
		return 0
	}
	startIndex := 0
	endIndex := 0
	maxLength := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			endIndex = i
		} else {
			if endIndex-startIndex+1 > maxLength {
				maxLength = endIndex - startIndex + 1
			}
			startIndex = i
		}
	}
	if endIndex-startIndex+1 > maxLength {
		maxLength = endIndex - startIndex + 1
	}
	return maxLength
}

func main() {
	fmt.Println(longestContinuousSubstring("abcde"))
}
