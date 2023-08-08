package main

import "fmt"

// tc: O(n^2)
// mc: O(1)

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

func main() {
	nums := []int{7, 2, 3, -4, 5, 6}
	fmt.Print(nums, "->")
	selectionSort(nums)
	fmt.Println(nums)
}
