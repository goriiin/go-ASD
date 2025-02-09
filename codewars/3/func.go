package main

import (
	"strings"
)

func LongestConsec(strArr []string, k int) string {
	if len(strArr) == 0 || k > len(strArr) || k <= 0 {
		return ""
	}

	b := strings.Builder{}

	if k == len(strArr) {
		for _, str := range strArr {
			b.WriteString(str)
		}

		return b.String()
	}

	lenArr := make([]int, len(strArr))

	for i := 0; i < len(strArr); i++ {
		lenArr[i] = len(strArr[i])
	}

	sum := 0

	for i := 0; i < k; i++ {
		sum += lenArr[i]
	}

	maxSum := sum
	index := 0

	for i := 1; i < len(lenArr)-k+1; i++ {
		sum -= lenArr[i-1]
		sum += lenArr[i+k-1]

		if sum > maxSum {

			index = i
			maxSum = sum
		}
	}

	for i := index; i < index+k; i++ {
		b.WriteString(strArr[i])
	}

	return b.String()
}
