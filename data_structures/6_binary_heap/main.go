package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

// Binary heap

// push: O(log n)
// pop: O(log n)
// depth: O(1)

type BinaryHeap[V constraints.Ordered] interface {
	Push(V)
	Pop()
	Root() (V, error)
	Depth() int
	Size() int
	Print()
}

type BHeap[V constraints.Ordered] struct {
	values []V
	size   int
}

func (bh *BHeap[V]) getParentIdx(idx int) int {
	if idx <= 0 {
		return -1
	}
	return (idx - 1) / 2
}

func (bh *BHeap[V]) getLeftChildIdx(idx int) int {
	res := 2*idx + 1
	if res <= 0 || res >= bh.size {
		return -1
	}
	return res
}

func (bh *BHeap[V]) getRightChildIdx(idx int) int {
	res := 2*idx + 2
	if res <= 0 || res >= bh.size {
		return -1
	}
	return res
}

func (bh *BHeap[V]) heapify(idx int) {
	leftIdx := bh.getLeftChildIdx(idx)
	rightIdx := bh.getRightChildIdx(idx)
	maxIdx := idx
	for {
		if leftIdx != -1 && bh.values[maxIdx] < bh.values[leftIdx] {
			maxIdx = leftIdx
		}
		if rightIdx != -1 && bh.values[maxIdx] < bh.values[rightIdx] {
			maxIdx = rightIdx
		}
		if maxIdx == idx {
			break
		}
		bh.values[idx], bh.values[maxIdx] = bh.values[maxIdx], bh.values[idx]
		leftIdx = bh.getLeftChildIdx(maxIdx)
		rightIdx = bh.getRightChildIdx(maxIdx)
		idx = maxIdx
	}
}

func (bh *BHeap[V]) printValue(idx int, spaceCount int) {
	if idx == -1 {
		return
	}
	const COUNT int = 2
	spaceCount += COUNT
	leftIdx := bh.getLeftChildIdx(idx)
	rightIdx := bh.getRightChildIdx(idx)
	bh.printValue(rightIdx, spaceCount)
	for i := COUNT; i < spaceCount; i++ {
		fmt.Print("   ")
	}
	fmt.Println(bh.values[idx])
	bh.printValue(leftIdx, spaceCount)
}

func (bh *BHeap[V]) Push(value V) {
	bh.values = append(bh.values, value)
	curIdx, parentIdx := bh.size, bh.getParentIdx(bh.size)
	for curIdx != 0 && bh.values[curIdx] > bh.values[parentIdx] {
		bh.values[curIdx], bh.values[parentIdx] = bh.values[parentIdx], bh.values[curIdx]
		curIdx, parentIdx = parentIdx, bh.getParentIdx(parentIdx)
	}
	bh.size++
}

func (bh *BHeap[V]) Pop() {
	if bh.size == 0 {
		return
	}
	bh.values[0] = bh.values[bh.size-1]
	bh.values = bh.values[: bh.size-1 : bh.size-1]
	bh.size--
	bh.heapify(0)
}

func (bh *BHeap[V]) Root() (V, error) {
	var res V
	if bh.size == 0 {
		return res, fmt.Errorf("empty heap")
	}
	return bh.values[0], nil
}

func (bh *BHeap[V]) Depth() int {
	if bh.size == 0 {
		return 0
	}
	return int(math.Log2(float64(bh.size))) + 1
}

func (bh *BHeap[V]) Size() int {
	return bh.size
}

func (bh *BHeap[V]) Print() {
	if bh.size == 0 {
		return
	}
	bh.printValue(0, 0)
}

func NewBinaryHeap[V constraints.Ordered]() BinaryHeap[V] {
	return &BHeap[V]{}
}

func main() {
	bh := NewBinaryHeap[int]()
	bh.Push(100)
	bh.Push(30)
	bh.Push(120)
	bh.Push(10)
	bh.Push(50)
	bh.Push(150)
	bh.Push(90)
	bh.Push(40)
	bh.Push(70)
	bh.Push(140)
	bh.Push(160)
	bh.Push(80)
	bh.Push(60)
	bh.Push(110)
	bh.Push(180)
	bh.Push(130)

	fmt.Printf("\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()

	bh.Pop()
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()

	bh.Pop()
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()

	bh.Pop()
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()

	bh.Pop()
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()

	bh.Pop()
	bh.Pop()
	bh.Pop()
	bh.Pop()
	bh.Pop()
	bh.Pop()
	bh.Pop()
	bh.Pop()
	bh.Pop()

	bh.Pop()
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()

	bh.Pop()
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()

	bh.Pop()
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bh.Size(), bh.Depth())
	bh.Print()
}
