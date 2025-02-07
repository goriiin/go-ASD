package main

func maxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var sum, maxSum int

	for i := 0; i < len(nums)-1; i++ {
		sum += nums[i]

		if nums[i] >= nums[i+1] {
			if sum > maxSum {
				maxSum = sum
			}

			sum = 0
		} else if i+2 == len(nums) {
			sum += nums[i+1]

			if sum > maxSum {
				maxSum = sum
			}

		}
	}

	return maxSum
}
