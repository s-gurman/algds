package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// tc: O(log n)
// mc: O(1)

func binarySearch[T constraints.Ordered](data []T, target T) int {
	pos := -1
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return mid
		}
		if data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return pos
}

func main() {
	arr := []int{14, 18, 19, 23, 24, 35}
	fmt.Println(binarySearch(arr, 18))
}
