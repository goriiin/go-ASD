package main

import (
	"bufio"
	"os"
)

func main() {
	var (
		reader  = bufio.NewReader(os.Stdin)
		printer = bufio.NewWriter(os.Stdout)
		ans     = "No"
		MFlag   = false
	)

	input, _, err := reader.ReadLine()
	if err != nil {
		return
	}

	str := string(input)

	for _, ch := range str {
		switch ch {
		case 'M':
			{
				MFlag = true
			}
		case 'R':
			{
				if !MFlag {
					ans = "Yes"
				}
				break
			}
		}
	}

	_, err = printer.WriteString(ans)
	if err != nil {
		return
	}

	err = printer.Flush()
	if err != nil {
		return
	}
}
