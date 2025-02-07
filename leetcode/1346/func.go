package main

import (
	"sort"
)

func checkIfExist(arr []int) bool {
	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		n := arr[i]

		j, ok := binSearch(arr, n*2)
		if !ok {
			continue
		}

		if i != j {
			return true
		}
	}

	return false
}

func binSearch(arr []int, target int) (int, bool) {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {
			return mid, true
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return 0, false
}
