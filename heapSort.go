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
	i := len(nums) - 1
	buildMaxHeap(nums)
	for i > 0 {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
		i--
	}
}

func buildMaxHeap(nums []int) {
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		heapify(nums, i, length)
	}

}
func heapify(nums []int, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	large := i
	if left < length && nums[large] < nums[left] {
		large = left
	}
	if right < length && nums[large] < nums[right] {
		large = right
	}
	if large != i {
		nums[i], nums[large] = nums[large], nums[i]
		heapify(nums, large, length)
	}
}
