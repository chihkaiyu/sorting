package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 4, 1, 3, 5},
	}

	for _, i := range input {
		fmt.Printf("Before: %v\n", i)
		sortedInput := bubbleSort(i)
		fmt.Printf("After: %v\n", sortedInput)
		fmt.Println()
	}
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}

	return nums
}
