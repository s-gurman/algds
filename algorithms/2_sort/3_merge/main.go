package main

import "fmt"

// tc: O(n log n)
// mc: O(n)

func mergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	leftNums := mergeSort(nums[:n/2])
	rightNums := mergeSort(nums[n/2:])

	il, ir := 0, 0
	nl, nr := len(leftNums), len(rightNums)
	result := make([]int, 0, nl+nr)

	for il < nl && ir < nr {
		if leftNums[il] < rightNums[ir] {
			result = append(result, leftNums[il])
			il++
		} else {
			result = append(result, rightNums[ir])
			ir++
		}
	}
	for il < nl {
		result = append(result, leftNums[il])
		il++
	}
	for ir < nr {
		result = append(result, rightNums[ir])
		ir++
	}

	return result
}

func main() {
	nums := []int{7, 2, 3, -4, 5, 6}
	fmt.Println(nums, "->", mergeSort(nums))
}
