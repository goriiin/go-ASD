package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	prefSum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefSum[i+1] = prefSum[i] + nums[i]
	}

	maxSumK := prefSum[k]

	for i := 0; i < len(prefSum)-k; i++ {
		sum := prefSum[i+k] - prefSum[i]
		if sum > maxSumK {
			maxSumK = sum
		}
	}

	return float64(maxSumK) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
