package main

import (
	"fmt"
	"reflect"
)

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 4, 1, 3, 5},
	}

	expected := []int{1, 2, 3, 4, 5}

	for _, i := range input {
		fmt.Printf("Before: %v\n", i)
		sortedInput := selectionSort(i)
		fmt.Printf("After: %v\n", sortedInput)
		fmt.Println()

		if !reflect.DeepEqual(sortedInput, expected) {
			fmt.Printf("Expected: %v\n", expected)
			fmt.Printf("Actual: %v\n", sortedInput)
			panic("Sorting error!")
		}
	}
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minimumIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minimumIndex] {
				nums[j], nums[minimumIndex] = nums[minimumIndex], nums[j]
			}
		}
	}

	return nums
}
