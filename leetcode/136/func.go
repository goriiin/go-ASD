package main

func singleNumber(nums []int) int {
	var ans int

	for _, v := range nums {
		ans ^= v
	}

	return ans
}
