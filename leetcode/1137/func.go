package main

func tribonacci(n int) int {
	trMap := make([]int, 38)
	trMap[0] = 0
	trMap[1] = 1
	trMap[2] = 1

	for i := 3; i <= n; i++ {
		trMap[i] = trMap[i-1] + trMap[i-2] + trMap[i-3]
	}

	return trMap[n]
}
