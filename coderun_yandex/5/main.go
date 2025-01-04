package main

import (
	"fmt"
)

func main() {
	var (
		memory       = make(map[int]uint64)
		n            int
		totalSeconds uint64
	)

	fmt.Scan(&n)

	if n == 1 {
		fmt.Println(1)

		return
	}

	if n == 2 {
		fmt.Println(2)

		return
	}

	memory[1] = 1
	memory[2] = 1
	totalSeconds = 2

	for i := 3; i <= n; i++ {
		memory[i] = memory[i-1] + memory[i-2]

		totalSeconds += memory[i]
	}

	fmt.Println(totalSeconds)
}
