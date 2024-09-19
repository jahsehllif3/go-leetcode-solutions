package main

import (
	"fmt"
	"sort"
)

// a. 反着来？最晚的时间 -> 最晚的车 -> （乘客时间唯一）最晚的能抢到车 ？有问题，因为还有公交车容量，导致前面的乘客不能上车，累积到后面
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {

	// b. 还是要得正着来，公交数量 * 公交容量 = 最大能上车的乘客数量（因为可能因为乘客时间太晚，公交没坐满）
	sort.Ints(buses)
	sort.Ints(passengers)
	busesNum := len(buses)
	theoryPassengersNum := busesNum * capacity

	minTime := 0
	maxTime := 0

	if len(passengers) > theoryPassengersNum {
		// 按乘客时间排序后，序号大于能上车的乘客数量先过滤掉
		//fmt.Println("过滤前乘客列表", passengers)
		passengers = passengers[0:theoryPassengersNum]
		//fmt.Println("过滤后乘客列表", passengers)

	}
	// 确定乘客时间后，其实理论最大时间也能确定，因为肯定是在最后一个乘客时间之前
	// 其实不太对，因为可能最后一辆公交没坐满，所以最大应该是最后一个公交时间
	maxTime = buses[busesNum-1]
	// b.1. 然后考虑，(乘客时间唯一,且公交容量有限)
	// 直接正向演绎
	passengerTimeMap := make(map[int]bool)
	busCapacityMap := make(map[int][]int)

	// 这里用双指针优化
	//for i := range buses {
	//	for j := range passengers {
	//		// 上过车了，跳过
	//		jPassenger := passengers[j]
	//		if passengerTimeMap[jPassenger] {
	//			continue
	//		}
	//		passengerTimeMap[jPassenger] = false
	//		// 公交车容量未满
	//		iBusPassengers := busCapacityMap[buses[i]]
	//		if iBusPassengers == nil {
	//			iBusPassengers = make([]int, 0)
	//		}
	//		if len(iBusPassengers) >= capacity {
	//			break
	//		}
	//		// 公交车时间大于等于乘客时间
	//		if buses[i] >= jPassenger {
	//			iBusPassengers = append(iBusPassengers, jPassenger)
	//			passengerTimeMap[jPassenger] = true
	//		}
	//		busCapacityMap[buses[i]] = iBusPassengers
	//	}
	//}
	passengerIndex := 0
	busCapacity := 0
	for i, _ := range buses {
		busCapacity = 0
		for passengerIndex < len(passengers) && busCapacity < capacity {
			jPassenger := passengers[passengerIndex]
			passengerTimeMap[jPassenger] = false
			// 公交车容量未满
			iBusPassengers := busCapacityMap[buses[i]]
			if iBusPassengers == nil {
				iBusPassengers = make([]int, 0)
			}
			// 公交车时间大于等于乘客时间
			if buses[i] >= jPassenger {
				iBusPassengers = append(iBusPassengers, jPassenger)
				passengerTimeMap[jPassenger] = true
				passengerIndex++
				busCapacity++
			} else {
				break
			}
			busCapacityMap[buses[i]] = iBusPassengers
		}
	}

	//fmt.Println(passengerTimeMap)
	//fmt.Println(busCapacityMap)

	// 公交时间 key 列表
	keys := make([]int, 0)
	for key := range busCapacityMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	//fmt.Println(keys)
	//
	//fmt.Println("遍历最小时间", minTime, "遍历最大时间", maxTime)
	for i := maxTime; i >= minTime; i-- {
		// 最后一辆没坐满，那么从最大索引倒着来即可
		if len(busCapacityMap[maxTime]) < capacity {
			if passengerTimeMap[i] != true {
				return i
			}
		} else {
			// 坐满了，那就从从最后一个人的时间往前倒着来，也就是抢座位
			x := busCapacityMap[maxTime]
			result := x[len(x)-1]
			if i < result && passengerTimeMap[i] != true {
				return i
			}
		}
	}
	return 999
}

func main() {
	fmt.Println(latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2))
	fmt.Println(latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2))
	fmt.Println(latestTimeCatchTheBus([]int{3}, []int{2, 4}, 2))
	fmt.Println(latestTimeCatchTheBus([]int{2}, []int{2}, 2))
	fmt.Println(latestTimeCatchTheBus([]int{3}, []int{4}, 1))
}
