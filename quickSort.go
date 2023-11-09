package main

import (
	"fmt"
	"math/rand"
)

// ascending
func main() {
	nums := []int{2, 3, 7, 9, 0, 4, 6, 17, 1, 5, 8, 2, 3, 7, 9, 0, 4, 6, 1, 5, 8}
	quickSort(nums)
	fmt.Print(nums)
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	randPivot(nums)
	p := partition(nums)
	quickSort(nums[0:p])
	quickSort(nums[p+1:])
}

func partition(nums []int) int {
	left := 0
	right := len(nums) - 1
	for right != left {
		for nums[right] >= nums[0] && right > left {
			right--
		}
		for nums[left] <= nums[0] && right > left {
			left++
		}
		swap(nums, left, right)
	}
	swap(nums, 0, right)
	return right
}

func randPivot(nums []int) {
	// fmt.Println("lenth: ", len(nums))
	i := rand.Intn(len(nums))
	swap(nums, 0, i)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}