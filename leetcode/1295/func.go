package main

import "math"

func findNumbers(nums []int) int {
	var count int

	for i := range nums {
		n := int(math.Log10(float64(nums[i]))) + 1

		if n%2 == 0 {
			count++
		}
	}

	return count
}
