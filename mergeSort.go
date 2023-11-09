package main

import "fmt"

func main() {
	var nums = []int{2, 5, 9, 1, 2, 5, 9, 1, 0, 7, 6, 3, 8, 4, 0, 7, 6, 3, 8, 4}
	sortedNums := mergeSort(nums)
	fmt.Print(sortedNums)

}

// ascending
func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	p := len(nums) / 2
	return merge(mergeSort(nums[0:p]), mergeSort(nums[p:]))
}
func merge(leftNums, rightNums []int) []int {
	var tempNums []int
	var i, j int
	for i < len(leftNums) && j < len(rightNums) {
		if leftNums[i] <= rightNums[j] {
			tempNums = append(tempNums, leftNums[i])
			i++
		} else {
			tempNums = append(tempNums, rightNums[j])
			j++
		}
	}
	if i < len(leftNums) {
		tempNums = append(tempNums, leftNums[i:]...)
	}
	if j < len(rightNums) {
		tempNums = append(tempNums, rightNums[j:]...)
	}

	return tempNums
}
