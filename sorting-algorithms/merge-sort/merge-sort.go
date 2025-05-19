package mergeSort

import "fmt"

// merge combines two sorted arrays into one sorted array
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	// Compare elements from both arrays and place the smaller one in the result
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	fmt.Printf("result: %v, left: %v, right %v \n\n", result, left, right)

	// Copy remaining elements from left array, if any
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	fmt.Printf("add to result from left - result = %v, left: %v \n\n", result, left)

	// Copy remaining elements from right array, if any
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	fmt.Printf("add to result from right - result = %v, right: %v \n\n", result, right)

	return result
}

func MergeSort(arr []int) []int {
	fmt.Printf("init: %v \n\n", arr)
	n := len(arr)
	// Create a copy to avoid modifying the original array
	result := make([]int, n)
	copy(result, arr)

	if n <= 1 {
		return result
	}

	// Divide the array into two halves
	mid := n / 2
	left := MergeSort(result[:mid])
	right := MergeSort(result[mid:])

	fmt.Printf("left: %v, right: %v \n\n", left, right)

	// Merge the sorted halves
	return merge(left, right)
}
