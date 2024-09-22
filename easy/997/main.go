package main

import (
	"fmt"
)

// trust[i] = [ai, bi] 表示 ai 信任 bi
// 任何人都信任法官，每个人相信的必须有交集
// 法官不信任任何人，即法官 ai 不出现
func findJudge(n int, trust [][]int) int {
	// n >= 1
	trustLen := len(trust)
	// n 个人，最多 n-1 个人信任法官，那么最大信任长度为 n*(n-1)
	maxTrustLen := n * (n - 1)
	if maxTrustLen < trustLen {
		return -1
	}
	/////////////////////////////////////////////
	// 演绎
	inDegree := make([]int, n+1)
	outDegree := make([]int, n+1)
	for _, v := range trust {
		inDegree[v[1]]++
		outDegree[v[0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegree[i] == n-1 && outDegree[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))                                 // 2
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))                 // -1
	fmt.Println(findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}})) // 3
	fmt.Println(findJudge(1, [][]int{}))                                       // 1
}
