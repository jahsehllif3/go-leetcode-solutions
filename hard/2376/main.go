package main

import (
	"fmt"
	"strconv"
)

// 看完题目马上想到遍历演绎，但是肯定超时。。
// 尝试下找规律，看到 leetcode 上提到动态规划，可惜不懂，先自己解一下
// 给定位数，非重复，排列组合，且数字小于等于 n
// 排列组合简单，但是卡在还要限定数字小于等于 n
// 还得是 dp
// 1. 状态转移方程 -> dp(i，mask,isLimit,isNum) =  Σ dp(pos + 1, newTight, newUsed)
// 即每一位
// 2. 记忆化
func countSpecialNumbers1(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	// 记忆化二维数组 [m][1 << 10]
	// 1 << 10 = 1024，即 0-9 十个数字 即 mask
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int, bool, bool) int
	// 当前位为 i，mask 表示 0-9 数字是否使用过，isLimit 表示是否受限，isNum 表示前面是否有数字
	dfs = func(i int, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			// int 默认值为 0
			return
		}
		if !isLimit && isNum {
			p := &memo[i][mask]
			if *p >= 0 { // 之前计算过
				return *p
			}
			// 避免在多个分支中重复赋值
			defer func() { *p = res }() // 记忆化
		}
		if !isNum { // 可以跳过当前数位
			res += dfs(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中，说明之前没有填过 d
				res += dfs(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return dfs(0, 0, true, false)
}
func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	// 记忆化二维数组 [m][1 << 10]，这里记忆条件只有 i 和 mask，因为 isLimit 和 isNum 几乎不会用到第二次
	// 1 << 10 = 1024，即 0-9 十个数字 即 mask
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int, bool, bool) int
	// 当前位为 i，mask 表示 0-9 数字是否使用过，isLimit 表示是否受限，isNum 表示前面是否有数字(前面输入的是否已经是有效数字)
	dfs = func(i int, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		// 无限制，且前面有数字，那可以从记忆里取
		if !isLimit && isNum {
			// 用指针
			p := &memo[i][mask]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		// 如果前面没有数字，那么可以跳过当前位(例如100，本次递归的只包含第一位的有效数字，所以可以加上跳过第一位的从第二位开始的递归)
		if !isNum {
			res += dfs(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if mask>>d&1 == 0 {
				res += dfs(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return dfs(0, 0, true, false)
}

func main() {
	fmt.Println(countSpecialNumbers(20))
	fmt.Println(countSpecialNumbers(5))
	fmt.Println(countSpecialNumbers(135))
}
