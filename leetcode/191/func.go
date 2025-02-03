package main

func hammingWeight(n int) int {
	var counter int

	for i := 0; i < 32; i++ {
		if 1<<i&n != 0 {
			counter++
		}
	}

	return counter
}
