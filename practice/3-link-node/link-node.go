package main

import "fmt"

type LinkNode struct {
	Data     int
	NextNode *LinkNode
}

// 在Go语言中，当你使用结构体的指针并访问它的字段时，Go会自动为你解引用
func main() {
	node0 := new(LinkNode)
	node0.Data = 0

	node1 := new(LinkNode)
	node1.Data = 1

	node2 := new(LinkNode)
	node2.Data = 2

	node0.NextNode = node1
	node1.NextNode = node2
	node2.NextNode = nil

	fmt.Println(node0)
	fmt.Println(&node0)
	fmt.Println(*node0)
	fmt.Println(*((*node0).NextNode))

	nowNode := node0
	for nowNode != nil {
		//fmt.Println(nowNode)
		//fmt.Println(nowNode.Data)
		nowNode = nowNode.NextNode
	}
}
