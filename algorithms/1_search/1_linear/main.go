package main

import "fmt"

// tc: O(n)
// mc: O(1)

func linearSearch[T comparable](data []T, target T) int {
	pos := -1
	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return i
		}
	}
	return pos
}

func main() {
	arr := []int{1, 3, 2, 6, 4}
	fmt.Println(linearSearch(arr, 6))
}
