package main

import "fmt"

// Stack (LIFO - Last In First Out)

// push: O(1)
// pop: O(1)

type Stack[V any] interface {
	Push(V)
	Pop()
	Head() (V, error)
	Size() int
}

type Node[V any] struct {
	next *Node[V]
	val  V
}

type LinkedStack[V any] struct {
	head *Node[V]
	size int
}

func (st *LinkedStack[V]) Push(value V) {
	node := new(Node[V])
	node.next = st.head
	node.val = value
	st.head = node
	st.size++
}

func (st *LinkedStack[V]) Pop() {
	if st.size == 0 {
		return
	}
	st.head = st.head.next
	st.size--
}

func (st *LinkedStack[V]) Head() (V, error) {
	var res V
	if st.size == 0 {
		return res, fmt.Errorf("empty stack")
	}
	return st.head.val, nil
}

func (st *LinkedStack[V]) Size() int {
	return st.size
}

func NewStack[V any]() Stack[V] {
	return &LinkedStack[V]{}
}

func main() {
	st := NewStack[int]()
	st.Push(5)
	st.Push(10)
	st.Push(15)

	for {
		size := st.Size()
		if size == 0 {
			fmt.Println("Stack is empty!")
			break
		}
		head, _ := st.Head()
		fmt.Printf("head = %v\nsize = %v\n\n", head, size)
		st.Pop()
	}

	head, found := st.Head()
	fmt.Printf("\nhead: %v\nfound: %s\n\n", head, found)
}
