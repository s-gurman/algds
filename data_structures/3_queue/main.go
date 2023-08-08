package main

import "fmt"

// Queue (FIFO - First In First Out)

// push: O(1)
// pop: O(1)

type Queue[V any] interface {
	Push(V)
	Pop()
	Front() (V, error)
	Back() (V, error)
	Size() int
}

type Node[V any] struct {
	next *Node[V]
	val  V
}

type LinkedQueue[V any] struct {
	front *Node[V]
	back  *Node[V]
	size  int
}

func (q *LinkedQueue[V]) Push(value V) {
	node := new(Node[V])
	node.next = nil
	node.val = value
	if q.size == 0 {
		q.front = node
	} else {
		q.back.next = node
	}
	q.back = node
	q.size++
}

func (q *LinkedQueue[V]) Pop() {
	if q.size <= 1 {
		q.front = nil
		q.back = nil
		q.size = 0
		return
	}
	q.front = q.front.next
	q.size--
}

func (q *LinkedQueue[V]) Front() (V, error) {
	var res V
	if q.size == 0 {
		return res, fmt.Errorf("empty queue")
	}
	return q.front.val, nil
}

func (q *LinkedQueue[V]) Back() (V, error) {
	var res V
	if q.size == 0 {
		return res, fmt.Errorf("empty queue")
	}
	return q.back.val, nil
}

func (q *LinkedQueue[V]) Size() int {
	return q.size
}

func NewQueue[V any]() Queue[V] {
	return &LinkedQueue[V]{}
}

func main() {
	q := NewQueue[int]()
	q.Push(5)
	q.Push(10)
	q.Push(15)

	for {
		size := q.Size()
		if size == 0 {
			fmt.Println("Queue is empty!")
			break
		}
		front, _ := q.Front()
		fmt.Printf("front = %v\nsize = %v\n\n", front, size)
		q.Pop()
	}

	front, found := q.Front()
	fmt.Printf("\nfront: %v\nfound: %s\n\n", front, found)
}
