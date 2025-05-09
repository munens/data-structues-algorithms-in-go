package sorting

import (
	"fmt"
)

//func MergeSortInPlace(arr []int) {
//	mid := int(math.Floor(float64(len(arr) / 2)))
//
//	fmt.Printf("init: arr = %v, mid = %d \n\n", arr, mid)
//
//	if len(arr) == 1 {
//		fmt.Printf("final arr = %v \n\n", arr)
//		return
//	}
//
//	MergeSortInPlace(arr[:mid])
//	MergeSortInPlace(arr[mid:])
//
//	fmt.Printf("apres: arr = %v \n\n", arr)
//}

// copilot attempt:

func mergeInPlace(arr []int, start, mid, end int) {
	// Temporary pointers
	i, j := start, mid+1

	fmt.Printf("mergeInPlace - start: arr = %v, i = %d, j = %d, start = %d,  mid = %d, end = %d \n\n", arr, i, j, start, mid, end)

	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			i++
		} else {
			// Element at j is smaller, shift elements to the right
			temp := arr[j]
			copy(arr[i+1:j+1], arr[i:j])
			arr[i] = temp

			// Update pointers
			i++
			mid++
			j++
		}
	}

	fmt.Printf("mergeInPlace - in-loop: arr = %v, i = %d, j = %d, start = %d,  mid = %d, end = %d \n\n", arr, i, j, start, mid, end)
}

func mergeSortInPlaceHelper(arr []int, start, end int) {
	if start >= end {
		fmt.Printf("return: arr = %v, start = %d, end = %d \n\n", arr, start, end)
		return
	}

	// Find the middle point
	mid := start + (end-start)/2

	fmt.Printf("arr = %v, start = %v, mid = %d, end = %d \n\n", arr, start, mid, end)

	// Recursively sort the two halves
	mergeSortInPlaceHelper(arr, start, mid)
	mergeSortInPlaceHelper(arr, mid+1, end)

	fmt.Printf("before merge: arr = %v, start = %d, mid = %d, end = %d \n\n", arr, start, mid, end)

	// Merge the sorted halves
	mergeInPlace(arr, start, mid, end)
}

func MergeSortInPlace(arr []int) {
	mergeSortInPlaceHelper(arr, 0, len(arr)-1)
	fmt.Printf("Sorted array: %v\n", arr)
}

// arr := []int{5, 3, 8, 4, 2}
