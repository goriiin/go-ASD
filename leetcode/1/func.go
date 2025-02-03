package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	var (
		num1, num2 int
	)

	copyNums := make([]int, len(nums))
	copy(copyNums, nums)

	sort.Ints(copyNums)
	l := 0

	fmt.Println(copyNums)
	for copyNums[l]+findMin(l, len(copyNums)-1) <= target || l <= len(copyNums)-2 {
		ind, ok := binarySearch(copyNums, target-copyNums[l])
		if ok {
			num1 = copyNums[ind]
			num2 = copyNums[l]

			break
		}

		l++
	}

	ans := make([]int, 2)
	var ok bool
	ans[0], ok = linearSearch(nums, findMin(num1, num2), -1)
	if !ok {
		return nil
	}

	ans[1], ok = linearSearch(nums, findMax(num1, num2), ans[0])
	if !ok {
		return nil
	}

	return ans
}

func linearSearch(nums []int, target, closed int) (int, bool) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && closed != i {
			return i, true
		}
	}

	return 0, false
}

func findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findMax(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func binarySearch(nums []int, target int) (int, bool) {
	var (
		l      = 0
		r      = len(nums) - 1
		middle = 0
		flag   = false
	)

	for l <= r {
		middle = (l + r) / 2

		if nums[middle] == target {
			flag = true
			break

		} else if nums[middle] > target {
			r = middle - 1
		} else if nums[middle] < target {
			l = middle + 1
		}
	}

	return middle, flag
}
