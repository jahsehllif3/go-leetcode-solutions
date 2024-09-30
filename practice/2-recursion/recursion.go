package main

import "fmt"

// 普通递归和尾部递归
// 尾递归版通过辅助函数和参数传递中间结果，减少了重复计算，同时具有更好的性能和空间使用效率
// 使用尾递归时，程序执行效率和内存占用都会更好，但并不是所有递归问题都适合尾递归
//
// 普通递归版
func fibonacci(n int) (res int) {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 尾部递归版
func fibonacciTail(n int) (res int) {
	return fibonacciTailCore(n, 0, 1)
}
func fibonacciTailCore(n int, a int, b int) (res int) {
	if n == 0 {
		return a
	}
	return fibonacciTailCore(n-1, b, a+b)
}
func main() {
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacciTail(10))
}
