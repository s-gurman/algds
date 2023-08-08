package main

import (
	"container/list"
	"encoding/json"
	"fmt"

	fasthash "github.com/segmentio/fasthash/fnv1a"
)

// Hash Table with Chaining Method

// get: O(1)
// insert: O(1)
// erase: O(1)

type HashTable[K comparable, V any] interface {
	Get(K) (V, bool)
	Insert(K, V)
	Erase(K)
	Size() int
}

type HTable[K comparable, V any] struct {
	hashFunc func([]byte) uint32
	buckets  []*list.List
	size     int
}

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func FindElement[K comparable, V any](l *list.List, key K) (_ *list.Element, val V, _ bool) {
	if l == nil || l.Len() == 0 {
		return nil, val, false
	}
	for p := l.Front(); p != nil; p = p.Next() {
		pair, _ := p.Value.(Pair[K, V])
		if pair.Key == key {
			return p, pair.Value, true
		}
	}
	return nil, val, false
}

func (ht *HTable[K, V]) getKeyIdx(key K) int {
	keyBytes, _ := json.Marshal(key)
	hash := ht.hashFunc(keyBytes)
	return int(hash % uint32(len(ht.buckets)))
}

func (ht *HTable[K, V]) loadFactor() float64 {
	bucketSize := len(ht.buckets)
	if bucketSize == 0 {
		return 1.
	}
	return float64(ht.size) / float64(bucketSize)
}

func (ht *HTable[K, V]) rehash() {
	newBuckets := make([]*list.List, 2*len(ht.buckets))
	for idx := 0; idx < len(ht.buckets); idx++ {
		l := ht.buckets[idx]
		if l == nil {
			continue
		}
		for p := l.Front(); p != nil; p = p.Next() {
			pair, _ := p.Value.(Pair[K, V])
			keyBytes, _ := json.Marshal(pair.Key)
			newIdx := int(ht.hashFunc(keyBytes) % uint32(len(newBuckets)))
			if newBuckets[newIdx] == nil {
				newBuckets[newIdx] = list.New()
			}
			newBuckets[newIdx].PushBack(pair)
		}
	}
	ht.buckets = newBuckets
}

func (ht *HTable[K, V]) Get(key K) (V, bool) {
	idx := ht.getKeyIdx(key)
	_, value, found := FindElement[K, V](ht.buckets[idx], key)
	return value, found
}

func (ht *HTable[K, V]) Insert(key K, val V) {
	if ht.loadFactor() > 0.7 {
		ht.rehash()
	}
	idx := ht.getKeyIdx(key)
	if ht.buckets[idx] == nil {
		ht.buckets[idx] = list.New()
	}
	elem, _, found := FindElement[K, V](ht.buckets[idx], key)
	if found {
		elem.Value = Pair[K, V]{key, val}
		return
	}
	ht.buckets[idx].PushBack(Pair[K, V]{key, val})
	ht.size++
}

func (ht *HTable[K, V]) Erase(key K) {
	idx := ht.getKeyIdx(key)
	elem, _, found := FindElement[K, V](ht.buckets[idx], key)
	if !found {
		return
	}
	ht.buckets[idx].Remove(elem)
	ht.size--
}

func (ht *HTable[K, V]) Size() int {
	return ht.size
}

func NewHashTable[K comparable, V any]() HashTable[K, V] {
	return &HTable[K, V]{
		hashFunc: fasthash.HashBytes32,
		buckets:  make([]*list.List, 2),
		size:     0,
	}
}

func main() {
	ht := NewHashTable[string, int]()
	ht.Insert("apple", 10)
	ht.Insert("apple", 1)
	ht.Insert("google", 2)
	ht.Insert("facebook", 3)
	ht.Insert("instagram", 4)
	ht.Insert("twitter", 5)
	ht.Insert("mailru", 6)
	ht.Insert("vk", 7)

	fmt.Printf("size = %v\n", ht.Size())

	num, found := ht.Get("apple")
	fmt.Println(num, found)

	num, found = ht.Get("instagram")
	fmt.Println(num, found)

	num, found = ht.Get("vk")
	fmt.Println(num, found)

	ht.Insert("apple", 10)
	num, found = ht.Get("apple")
	fmt.Println(num, found)

	ht.Erase("apple")
	num, found = ht.Get("apple")
	fmt.Println(num, found)

	ht.Erase("google")
	ht.Erase("facebook")
	ht.Erase("instagram")
	ht.Erase("twitter")
	ht.Erase("mailru")
	ht.Erase("vk")

	// for _, bucket := range ht.buckets {
	// 	if bucket != nil {
	// 		fmt.Print(bucket.Len(), " ")
	// 	}
	// }
	// fmt.Println()
}
