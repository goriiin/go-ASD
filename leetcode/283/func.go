package main

func moveZeroes(nums []int) {
	var j, zeroCount int
	var flag bool

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			if !flag {
				j = i
				flag = true
			}
		}

		if nums[i] != 0 {
			if flag {
				nums[j] = nums[i]
			}
			j++
		}
	}

	for i := len(nums) - 1; i >= len(nums)-zeroCount; i-- {
		nums[i] = 0
	}
}
