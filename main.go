package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nodeData := []int{1, 2, 6, 3, 4, 5, 6, 1}

	virtualNode := &ListNode{}

	tmp := virtualNode
	for _, datum := range nodeData {
		tmp.Val = datum
		newness := new(ListNode)
		tmp.Next = newness
		tmp = tmp.Next
	}

	fmt.Printf("%#v", virtualNode)
}
