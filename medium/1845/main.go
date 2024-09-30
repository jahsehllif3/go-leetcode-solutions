package main

import (
	"container/heap"
	"sort"
)

type SeatManager struct {
	sort.IntSlice // 继承 Len, Less, Swap
}

func Constructor(n int) SeatManager {
	m := SeatManager{make([]int, n)}
	for i := range m.IntSlice {
		m.IntSlice[i] = i + 1
	}
	// 有序数组无需堆化
	return m
}
func (m *SeatManager) Reserve() int {
	return heap.Pop(m).(int)
}

func (m *SeatManager) Unreserve(seatNumber int) {
	heap.Push(m, seatNumber)
}

func (m *SeatManager) Push(v any) { m.IntSlice = append(m.IntSlice, v.(int)) }
func (m *SeatManager) Pop() any {
	a := m.IntSlice
	v := a[len(a)-1]
	m.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
//func main() {
//	obj := Constructor(n)
//	param_1 := obj.Reserve()
//	obj.Unreserve(seatNumber)
//	// [-1,-1,-1,5,4,4,-1,-1,-1]
//	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
//}
