package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	var (
		in    = bufio.NewReader(os.Stdin)
		out   = bufio.NewWriter(os.Stdout)
		ans   = "No"
		MFlag = false
	)

	defer func(printer *bufio.Writer) {
		err := printer.Flush()
		if err != nil {
			log.Fatalln(err)
		}
	}(out)

	input, _, err := in.ReadLine()
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

	_, err = out.WriteString(ans)
	if err != nil {
		return
	}
}
