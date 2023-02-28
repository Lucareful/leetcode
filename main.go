package main

// Node 单链表
type Node struct {
	Data int
	Next *Node
}

// Nodes 双链表
type Nodes struct {
	Prev *Node
	Data int
	Next *Node
}
