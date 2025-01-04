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
		err        error
		reader     = bufio.NewReader(os.Stdin)
		prev, this int
		first      = true
	)

	input, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	for _, num := range strings.Split(strings.TrimSpace(input), " ") {
		this, err = strconv.Atoi(num)
		if first {
			prev = this
			first = false

			continue
		}

		if prev >= this {
			fmt.Println("NO")

			return
		}
		prev = this

	}

	fmt.Println("YES")

}
