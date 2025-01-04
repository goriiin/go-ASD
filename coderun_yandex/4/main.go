package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		array  = make([]int, 0, 16)
	)

	input, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	for _, num := range strings.Split(strings.TrimSpace(input), " ") {
		n, err := strconv.Atoi(num)
		if err != nil {
			return
		}

		array = append(array, n)
	}

	counter := 0

	for i := 1; i < len(array)-1; i++ {
		if array[i-1] < array[i] && array[i] > array[i+1] {
			counter++
		}
	}

	fmt.Println(counter)

}
