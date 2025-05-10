package sorting

import "fmt"

// []int{5, 3, 8, 4, 2}
// []int{5, 3, 4, 2, /8}
// []int{5, 3, 4, 2, /8}

func QuickSort(arr []int) []int {
	fmt.Printf("init: %v \n\n", arr)

	if len(arr) <= 1 {
		return arr
	}

	pivotIdx := len(arr) / 2
	pivot := arr[pivotIdx]

	fmt.Printf("pivotIdx: %d, pivot = %d \n\n", pivotIdx, pivot)

	newArr := make([]int, len(arr))

	idxBeforePivot := 0
	newPivotIdx := 0
	idxAfterPivot := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		vi := arr[i]

		if vi == pivot {
			continue
		}

		if vi <= pivot {
			newArr[idxBeforePivot] = vi
			idxBeforePivot++
			newPivotIdx++
		} else {
			newArr[idxAfterPivot] = vi
			idxAfterPivot--
		}

		fmt.Printf("newArr = %v, idxBeforePivot = %d, idxAfterPivot = %d, newPivotIdx = %d \n", newArr, idxBeforePivot, idxAfterPivot, newPivotIdx)
	}

	newArr[newPivotIdx] = pivot

	fmt.Printf("newArr = %v, \n\n", newArr)

	left := QuickSort(newArr[0:newPivotIdx])
	right := QuickSort(newArr[newPivotIdx+1 : len(newArr)])

	fmt.Printf("left = %v, pivot = %d, right = %v \n", left, pivot, right)

	final := append(append(left, pivot), right...)

	fmt.Printf("final = %v \n", final)

	return final
}

// 6/5 - there is a problem in the algorithm in terms of how we place values before and
//       after pivot. This requires further analysis.
// 6/8 - it turns out that the append function on line 56 and 58 was increasing the size of the array, newArr
//       This is why the lenght and capacity kept increasing. Not sure why the append was doing this and when
//       append does this sort of thing vs. when it doesnt. Need to find out
