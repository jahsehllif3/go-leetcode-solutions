package main

import "fmt"

// https://www.bilibili.com/video/BV1AD4y1k7pb
//
// 汉诺塔，每次只能移动一个盘子，且大盘子不能放在小盘子上
//
// 1 个盘子，直接从 a 移动到 c（只有1个盘，不需要借助） （1,a,b,c）
//
// 2 个盘子
// 盘1 从 a 移动到 b（只有1个盘，不需要借助）(1,a,c,b)
// 盘2 从 a 移动到 c (1,a,b,c)
// 盘1 从 b 移动到 c（只有1个盘，不需要借助） (1,b,a,c)
//
// 3 个盘子
// 盘1 从 a 移动到 c 盘2 从 a 移动到 b 盘1 从 c 移动到 b (将盘1、盘2，两个盘，移动到 b，借助 c) (2,a,c,b)
// 盘3 从 a 移动到 c (1,a,b,c)
// 盘1 从 b 移动到 a 盘2 从 b 移动到 c 盘1 从 a 移动到 c（将盘1、盘2，连个盘，移动到 c，借助 a） (2,b,a,c)
//
// 借助的原因是，大盘子不能放在小盘子上
//
// 将 n 个盘子 从 a 移动到 c (n,a,b,c)
// 1.将 n-1 个盘子从 a 移动到 b（借助 c） (n-1,a,c,b)
// 2.将第 n 个盘子从 a 移动到 c (1,a,b,c)
// 3.将 n-1 个盘子从 b 移动到 c（借助 a） (n-1,b,a,c)
// 移动次数=2^n-1
//
//	func main() {
//		move(3, "A", "B", "C")
//	}
//
// // n 个盘子，从 a 移动到 c，借助 b
//
//	func move(n int, a, b, c string) {
//		if n == 1 {
//			fmt.Println(a, "->", c)
//			return
//		}
//		move(n-1, a, c, b)
//		fmt.Println(a, "->", c)
//		move(n-1, b, a, c)
//	}
//
// n 个盘子，从 a 移动到 c，借助 b
func move(n int, a, b, c string, towers map[string][]int) {
	fmt.Println("Move", n, "from", a, "to", c)
	if n == 1 {
		// 从柱子 a 移动到柱子 c
		moveDisk(a, c, towers)
		printTowers(towers)
		return
	}
	move(n-1, a, c, b, towers) // 将 n-1 个盘子从 a 移到 b，借助 c
	moveDisk(a, c, towers)     // 将第 n 个盘子从 a 移到 c
	printTowers(towers)        // 输出当前塔的状态
	move(n-1, b, a, c, towers) // 将 n-1 个盘子从 b 移到 c，借助 a
}

// 移动盘子
func moveDisk(from, to string, towers map[string][]int) {
	// 从 from 移动一个盘子到 to
	disk := towers[from][len(towers[from])-1]         // 获取 from 顶部的盘子
	towers[from] = towers[from][:len(towers[from])-1] // 移除 from 顶部的盘子
	towers[to] = append(towers[to], disk)             // 将盘子放到 to
}

// 输出当前柱子的状态
func printTowers(towers map[string][]int) {
	fmt.Println("A:", towers["A"])
	fmt.Println("B:", towers["B"])
	fmt.Println("C:", towers["C"])
	fmt.Println("------------")
}

func main() {
	// 定义塔的初始状态
	n := 4
	// 初始化塔的状态，A 上有所有的盘子，B 和 C 为空
	towers := map[string][]int{
		"A": make([]int, n),
		"B": []int{},
		"C": []int{},
	}

	// 初始化 A 塔上的盘子，盘子编号为 n 到 1
	for i := 0; i < n; i++ {
		towers["A"][i] = n - i
	}

	// 打印初始状态
	fmt.Println("Initial state:")
	printTowers(towers)

	// 开始移动盘子
	move(n, "A", "B", "C", towers)
}
