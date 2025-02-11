package main

import "math"

func Gap(g, m, n int) []int {
	first, second := m, g+m
	if second > n {
		return []int{0, 0}
	}

	for second <= n {
		if !search(first) && !search(second) {
			return []int{first, second}
		}

		first++
		second++
	}

	return []int{0, 0}
}

func search(target int) bool {
	if target < 4 {
		return false
	}

	l, r := 2, int(math.Sqrt(float64(target)))+1

	for l <= r {
		if target%l == 0 {
			return true
		}

		l++
	}

	return false
}
