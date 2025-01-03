package main

import (
	"sort"
)

func avgV2(a, b, c int) int {
	nums := []int{a, b, c}
	sort.Ints(nums)

	return nums[1]
}
