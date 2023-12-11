// https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150

package main

func mergeSort(nums1 []int, m int, nums2 []int, n int) []int {
	var fin []int
	var prim []int
	var sec []int
	if nums1[0] <= nums2[0] {
		prim = nums1
		sec = nums2
	} else {
		prim = nums2
		sec = nums1
	}

	secMap := make(map[int][]int)

	for i := 0; i < len(sec); i++ {
		v := sec[i]
		_, ok := secMap[v]
		if ok {
			secMap[v] = append(secMap[v], v)
		} else {
			secMap[v] = []int{v}
		}
	}

	i := 0

	for {

		var primV int
		if i < m {
			primV = prim[i]
			fin = append(fin, primV)
		} else {
			primV = i
		}

		if v, ok := secMap[primV]; ok {
			fin = append(fin, v...)
		}

		if len(fin) == (m + n) {
			break
		} else {
			i++
		}
	}

	return fin
}

func main() {
	mergeSort([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 2, 5, 6, 6}, 5)
}
