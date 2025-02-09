package main

func FindNb(m int) int {
	i := 1
	var sum int

	for {
		newSum := sum + cube(i)
		if sum < m && m < newSum {
			return -1
		} else if newSum == m {
			return i
		}

		sum = newSum
		i++
	}
}

func cube(n int) int {
	return n * n * n
}
