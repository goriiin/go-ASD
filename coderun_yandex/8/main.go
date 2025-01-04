package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		reader   = bufio.NewReader(os.Stdin)
		array    = make([]string, 0, 10)
		wordsMap = make(map[string]struct{})
	)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		array = append(array, strings.Fields(line)...)
	}

	for _, c := range array {
		wordsMap[c] = struct{}{}
	}

	fmt.Println(len(wordsMap))
}
