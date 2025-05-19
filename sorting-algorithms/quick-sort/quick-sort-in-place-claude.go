package quickSort

import "fmt"

// partition rearranges the array and returns the pivot index
func partition(arr []int, low, high int) int {
	// Choose the rightmost element as the pivot
	pivot := arr[high]

	// Index of the smaller element
	i := low - 1

	for j := low; j < high; j++ {
		fmt.Printf("partition-loop-start: arr = %v, i = %d, j = %d, arr[j] = %d,  pivot = %d \n", arr, i, j, arr[j], pivot)
		// If the current element is smaller than or equal to the pivot
		if arr[j] <= pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}

		fmt.Printf("  partition-loop-end: arr = %v, i = %d, j = %d, pivot = %d \n\n", arr, i, j, pivot)
	}

	// Swap arr[i+1] and arr[high] (pivot)
	arr[i+1], arr[high] = arr[high], arr[i+1]

	fmt.Printf("partition: arr = %v, pivot = %d \n\n", arr, pivot)

	return i + 1
}

// quickSortHelper is a recursive helper function for quick sort
func quickSortHelper(arr []int, low, high int) {
	fmt.Printf("quickSortHelper-start: arr = %v, low = %d, high = %d \n\n", arr, low, high)
	if low < high {
		// Partition the array and get the pivot index
		pivotIdx := partition(arr, low, high)
		fmt.Printf("quickSortHelper: pivotIdx = %d \n\n", pivotIdx)

		// Recursively sort the subarrays
		quickSortHelper(arr, low, pivotIdx-1)
		quickSortHelper(arr, pivotIdx+1, high)
	}
}

func QuickSortInPlaceClaude(arr []int) []int {

	fmt.Printf("init: %v \n\n", arr)
	n := len(arr)

	quickSortHelper(arr, 0, n-1)
	return arr
}
