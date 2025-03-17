package sorting

import (
	"fmt"
	"math"
)

func selectionSort(arr []int, idx int) {

	fmt.Printf("initial arr = %v, idx = %d \n\n", arr, idx)

	smallest := int(math.Inf(1))
	smallestIdx := 0
	for i := idx; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIdx = i
		}
	}

	v0 := arr[idx]
	v1 := arr[smallestIdx]
	arr[idx] = v1
	arr[smallestIdx] = v0

	fmt.Printf("final arr = %v, smallest = %d, smallestIdx = %d \n\n", arr, smallest, smallestIdx)

	if idx == len(arr)-1 {
		return
	} else {
		selectionSort(arr, idx+1)
	}
}

func SelectionSort(arr []int) {
	selectionSort(arr, 0)
}
