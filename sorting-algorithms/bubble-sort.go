package sorting

import "fmt"

func BubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	swaps := 0

	fmt.Printf("initial arr = %v \n\n", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			v0 := arr[i-1]
			v1 := arr[i]
			arr[i] = v0
			arr[i-1] = v1
			swaps++
		}
	}

	fmt.Printf("final arr = %v, swaps = %d \n\n", arr, swaps)

	if swaps == 0 {
		return
	} else {
		BubbleSort(arr)
	}
}
