package main

import (
	"fmt"
)

// ascending
func main() {
	nums := []int{2, 3, 7, 9, 0, 4, 6, 17, 1, 5, 8, 2, 3, 7, 9, 0, 4, 6, 1, 5, 8}
	heapSort(nums)
	fmt.Print(nums)
}

func heapSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		sortedNums := nums[i:]
		heapify(sortedNums)
	}
}

func heapify(nums []int) {
	n := len(nums) - 1
	for i := n/2 - 1; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		shiftUp(nums, i, left, right)
	}
	shiftUp(nums, 0, 0, n)
}

func shiftUp(nums []int, i, left, right int) {
	if nums[i] > nums[right] {
		nums[i], nums[right] = nums[right], nums[i]
	}
	if nums[i] > nums[left] {
		nums[i], nums[left] = nums[left], nums[i]
	}
}