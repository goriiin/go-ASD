package main

func MaximumSubarraySum(numbers []int) int {
	prefSum := make([]int, len(numbers)+1)

	for i := 1; i < len(numbers)+1; i++ {
		prefSum[i] += prefSum[i-1] + numbers[i-1]
	}

	maxNum := prefSum[0]
	minNum := prefSum[0]

	for _, i := range prefSum {
		if i > maxNum {
			maxNum = i
		}
		if i < minNum {
			minNum = i
		}
	}

	return maxNum - minNum
}
