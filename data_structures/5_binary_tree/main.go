package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Binary search tree without balancing

// add: O(log n)
// remove: O(log n)
// contain: O(log n)
// depth: O(log n)

type BinaryTree[V constraints.Ordered] interface {
	Add(V)
	Remove(V)
	IsContain(V) bool
	Depth() int
	Size() int
	Print()
}

type Node[V constraints.Ordered] struct {
	left  *Node[V]
	right *Node[V]
	val   V
}

type BTree[V constraints.Ordered] struct {
	root *Node[V]
	size int
}

func (bt *BTree[V]) addNode(node *Node[V], value V) {
	if value < node.val && node.left == nil {
		newNode := new(Node[V])
		newNode.left = nil
		newNode.right = nil
		newNode.val = value
		node.left = newNode
	} else if value < node.val {
		bt.addNode(node.left, value)
	} else if node.right == nil {
		newNode := new(Node[V])
		newNode.left = nil
		newNode.right = nil
		newNode.val = value
		node.right = newNode
	} else {
		bt.addNode(node.right, value)
	}
}

func (bt *BTree[V]) removeNode(node *Node[V], parent *Node[V]) {
	if node.left == nil && node.right == nil {
		if parent == nil {
			bt.root = nil
			return
		}
		if node.val < parent.val {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}
	if node.right == nil {
		node.val = node.left.val
		node.right = node.left.right
		node.left = node.left.left
		return
	}
	if node.left == nil {
		node.val = node.right.val
		node.left = node.right.left
		node.right = node.right.right
		return
	}
	parent = node
	minNode := node.right
	for minNode.left != nil {
		parent = minNode
		minNode = minNode.left
	}
	node.val = minNode.val
	bt.removeNode(minNode, parent)
}

func (bt *BTree[V]) findNode(value V) (node *Node[V], parent *Node[V]) {
	for node = bt.root; node != nil; {
		if value < node.val {
			parent = node
			node = node.left
			continue
		}
		if value > node.val {
			parent = node
			node = node.right
			continue
		}
		break
	}
	return
}

func (bt *BTree[V]) depthNode(node *Node[V]) int {
	if node == nil {
		return 0
	}
	leftMax := bt.depthNode(node.left)
	rightMax := bt.depthNode(node.right)
	if leftMax < rightMax {
		return rightMax + 1
	}
	return leftMax + 1
}

func (bt *BTree[V]) dumpNode(node *Node[V], spaceCount int) {
	if node == nil {
		return
	}
	const COUNT int = 2
	spaceCount += COUNT
	bt.dumpNode(node.right, spaceCount)
	for i := COUNT; i < spaceCount; i++ {
		fmt.Print("   ")
	}
	fmt.Println(node.val)
	bt.dumpNode(node.left, spaceCount)
}

func (bt *BTree[V]) Add(value V) {
	if bt.size == 0 {
		bt.root = new(Node[V])
		bt.root.left = nil
		bt.root.right = nil
		bt.root.val = value
	} else {
		bt.addNode(bt.root, value)
	}
	bt.size++
}

func (bt *BTree[V]) Remove(value V) {
	node, parent := bt.findNode(value)
	if node == nil {
		return
	}
	bt.removeNode(node, parent)
	bt.size--
}

func (bt *BTree[V]) IsContain(value V) bool {
	if bt.size == 0 {
		return false
	}
	node, _ := bt.findNode(value)
	return node != nil
}

func (bt *BTree[V]) Depth() int {
	return bt.depthNode(bt.root)
}

func (bt *BTree[V]) Size() int {
	return bt.size
}

func (bt *BTree[V]) Print() {
	bt.dumpNode(bt.root, 0)
}

func NewBinaryTree[V constraints.Ordered]() BinaryTree[V] {
	return &BTree[V]{}
}

func main() {
	bt := NewBinaryTree[int]()
	bt.Add(100)
	bt.Add(30)
	bt.Add(120)
	bt.Add(10)
	bt.Add(50)
	bt.Add(150)
	bt.Add(0)
	bt.Add(40)
	bt.Add(70)
	bt.Add(130)
	bt.Add(160)
	bt.Add(45)
	bt.Add(42)
	bt.Add(47)

	fmt.Printf("\nsize: %d\ndepth: %d\n", bt.Size(), bt.Depth())
	fmt.Println(bt.IsContain(70), bt.IsContain(130))
	bt.Print()

	bt.Remove(100)
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bt.Size(), bt.Depth())
	fmt.Println(bt.IsContain(100))
	bt.Print()

	bt.Remove(42)
	bt.Remove(47)
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bt.Size(), bt.Depth())
	fmt.Println(bt.IsContain(42), bt.IsContain(47))
	bt.Print()

	bt.Remove(30)
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bt.Size(), bt.Depth())
	fmt.Println(bt.IsContain(30))
	bt.Print()

	bt.Remove(120)
	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bt.Size(), bt.Depth())
	fmt.Println(bt.IsContain(120))
	bt.Print()

	bt.Remove(10)
	bt.Remove(50)
	bt.Remove(150)
	bt.Remove(0)
	bt.Remove(40)
	bt.Remove(70)
	bt.Remove(130)
	bt.Remove(160)
	bt.Remove(45)

	fmt.Printf("\n\nsize: %d\ndepth: %d\n", bt.Size(), bt.Depth())
	bt.Print()
}
