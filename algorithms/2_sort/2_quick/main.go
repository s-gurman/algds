package main

import (
	"fmt"
	"math/rand"
)

// tc: O(n log n)
// mc: O(n)

func quickSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	lessNums := make([]int, 0, n)
	greatNums := make([]int, 0, n)
	pivot := nums[rand.Intn(n)]

	for i := 0; i < n; i++ {
		if nums[i] < pivot {
			lessNums = append(lessNums, nums[i])
		} else {
			greatNums = append(greatNums, nums[i])
		}
	}

	lessNums = quickSort(lessNums)
	greatNums = quickSort(greatNums)

	return append(lessNums, greatNums...)
}

func main() {
	nums := []int{7, 2, 3, -4, 5, 6}
	fmt.Println(nums, "->", quickSort(nums))
}
