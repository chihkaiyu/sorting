package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 4, 1, 3, 5},
	}

	expected := []int{1, 2, 3, 4, 5}

	for _, i := range test {
		fmt.Printf("Before: %v\n", i)
		buildMinHeap(i)
		sortedHeap := heapSort(i)
		fmt.Printf("After: %v\n", sortedHeap)
		fmt.Println()

		if !reflect.DeepEqual(sortedHeap, expected) {
			fmt.Printf("Expected: %v\n", expected)
			fmt.Printf("Actual: %v\n", sortedHeap)
			panic("Sorting error!")
		}
	}

}

func minHeapify(heap []int, parent int) {
	minimumChild := parent*2 + 1
	if minimumChild >= len(heap) {
		return
	}

	right := minimumChild + 1

	if right < len(heap) && heap[right] < heap[minimumChild] {
		minimumChild++
	}

	if heap[minimumChild] < heap[parent] {
		heap[minimumChild], heap[parent] = heap[parent], heap[minimumChild]

		minHeapify(heap, minimumChild)
	}
}

func buildMinHeap(heap []int) {
	lastParent := (len(heap) - 2) / 2

	for i := lastParent; i > -1; i-- {
		minHeapify(heap, i)
	}
}

func heapSort(heap []int) []int {
	var sortedHeap []int

	for len(heap) > 1 {
		heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
		sortedHeap = append(sortedHeap, heap[len(heap)-1])
		heap = heap[:len(heap)-1]

		minHeapify(heap, 0)
	}

	sortedHeap = append(sortedHeap, heap[0])

	return sortedHeap
}
