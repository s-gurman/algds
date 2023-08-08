package main

import "fmt"

// Double-linked list

// push_back: O(1)
// push_front: O(1)
// pop_back: O(1)
// pop_front: O(1)

type LinkedList[V any] interface {
	PushBack(V)
	PushFront(V)
	PopBack()
	PopFront()
	Front() (V, error)
	Back() (V, error)
	Size() int
}

type Node[V any] struct {
	next *Node[V]
	prev *Node[V]
	val  V
}

type List[V any] struct {
	front *Node[V]
	back  *Node[V]
	size  int
}

func (l *List[V]) PushBack(value V) {
	node := new(Node[V])
	node.next = nil
	node.prev = l.back
	node.val = value
	if l.size == 0 {
		l.front = node
	} else {
		l.back.next = node
	}
	l.back = node
	l.size++
}

func (l *List[V]) PushFront(value V) {
	node := new(Node[V])
	node.prev = nil
	node.next = l.front
	node.val = value
	if l.size == 0 {
		l.back = node
	} else {
		l.front.prev = node
	}
	l.front = node
	l.size++
}

func (l *List[V]) PopBack() {
	if l.size <= 1 {
		l.front = nil
		l.back = nil
		l.size = 0
		return
	}
	l.back.prev.next = nil
	l.back = l.back.prev
	l.size--
}

func (l *List[V]) PopFront() {
	if l.size <= 1 {
		l.front = nil
		l.back = nil
		l.size = 0
		return
	}
	l.front.next.prev = nil
	l.front = l.front.next
	l.size--
}

func (l *List[V]) Front() (V, error) {
	var res V
	if l.size == 0 {
		return res, fmt.Errorf("empty list")
	}
	return l.front.val, nil
}

func (l *List[V]) Back() (V, error) {
	var res V
	if l.size == 0 {
		return res, fmt.Errorf("empty list")
	}
	return l.back.val, nil
}

func (l *List[V]) Size() int {
	return l.size
}

func NewList[V any]() LinkedList[V] {
	return &List[V]{}
}

func main() {
	list := NewList[int]()
	list.PushBack(72)
	list.PushFront(-21)
	list.PushBack(43)
	list.PushFront(17)
	list.PushBack(-16)

	for {
		size := list.Size()
		if size == 0 {
			fmt.Println("List is empty!")
			break
		}
		front, _ := list.Front()
		back, _ := list.Back()
		fmt.Printf("front = %v back = %v\nsize = %v\n\n", front, back, size)
		if size%2 == 0 {
			list.PopBack()
		} else {
			list.PopFront()
		}
	}

	front, found := list.Front()
	fmt.Printf("\nfront: %v\nfound: %s\n", front, found)

	back, found := list.Back()
	fmt.Printf("\nback: %v\nfound: %s\n\n", back, found)
}
