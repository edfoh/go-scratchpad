package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

You must solve it in O(n) time complexity.
*/

func FindKthLargest(nums []int, k int) int {
	return kLargest(nums, 0, len(nums)-1, k-1)
}

func findPartition(nums []int, low, high int) int {
	pivot := high

	pivotVal := nums[pivot]

	j := low

	for i := low; i < pivot; i++ {
		if nums[i] > pivotVal {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	nums[j], nums[pivot] = nums[pivot], nums[j]
	return j
}

func kLargest(nums []int, low, high, k int) int {
	pi := findPartition(nums, low, high)

	if pi == k {
		return nums[pi]
	}

	if pi > k {
		return kLargest(nums, low, pi-1, k)
	}

	return kLargest(nums, pi+1, high, k)

}
