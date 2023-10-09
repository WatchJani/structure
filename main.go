package main

import (
	"fmt"
	"time"
)

type Node struct {
	value int
	node  *Node
}

type Tree struct {
	LastNode *Node
	Memory   []*Node
}

func NewStruct(capacity int) Tree {
	return Tree{
		Memory: make([]*Node, 0, capacity),
	}
}

func New(value int) *Node {
	return &Node{
		value: value,
	}
}

func (t *Tree) Inset(value int) {
	newNode := New(value)

	if t.LastNode != nil {
		t.LastNode.node = newNode
	}

	t.LastNode = newNode
	t.Memory = append(t.Memory, newNode)
}

func (t *Tree) Read() {
	for _, value := range t.Memory {
		fmt.Println(value.value)
	}
}

func main() {
	structure := NewStruct(1000)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 90, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 90, 1, 2, 3, 4, 5, 6, 7, 8, 90, 1, 2, 3, 4, 5, 6, 7, 8, 90, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	start := time.Now()
	for _, number := range numbers {
		structure.Inset(number)
	}
	fmt.Println(time.Since(start))

	structure.Read()
}
