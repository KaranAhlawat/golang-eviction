package model

// Node A simple node in a doubly linked list.
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}
