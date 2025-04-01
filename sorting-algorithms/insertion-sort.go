package sorting

import "fmt"

/*
attempt 1: - this attempt was done wrongly. See this link - https://miro.com/app/board/uXjVIS3ll7M=/?moveToWidget=3458764622254828741&cot=14

func insertionSort(arr []int, idx int) {
	smallest := arr[idx]

	fmt.Printf("arr = %v, idx = %d \n\n", arr, idx)

	for i := idx + 1; i < len(arr); i++ {
		fmt.Printf("init: arr = %v, i = %d, idx = %d, smallest = %d, arr[i] = %d \n\n", arr, i, idx, smallest, arr[i])
		if arr[i] < smallest {
			vi := arr[i]
			arr[idx] = vi
			arr[i] = smallest
			smallest = vi
		}
		fmt.Printf("next: arr = %v, idx = %d \n\n", arr, idx)
	}

	if idx == len(arr)-1 {
		return
	} else {
		insertionSort(arr, idx+1)
	}
}

*/

func insertionSort(arr []int, idx int) {
	fmt.Printf("arr = %v, idx = %d \n\n", arr, idx)
	curr := arr[idx]
	for i := idx - 1; i >= 0; i-- {
		fmt.Printf("init: arr = %v, i = %d, idx = %d, curr = %d, arr[i] = %d \n\n", arr, i, idx, curr, arr[i])
		if arr[i] > curr {
			vi := arr[i]
			arr[i] = curr
			arr[i+1] = vi
		}
	}

	if idx == len(arr)-1 {
		return
	} else {
		insertionSort(arr, idx+1)
	}
}

func InsertionSort(arr []int) {
	insertionSort(arr, 1)
}
