package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getKey(s string, start int) string {
	var sb strings.Builder
	sb.Grow((len(s) + 1) / 2)
	for i := start; i < len(s); i += 2 {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func wordsCalc(words map[string]int) int {
	sum := 0

	for key, m := range words {
		if len(key) > 0 {
			sum += m * (m - 1) / 2
		}
	}

	return sum
}

func main() {
	var (
		in  = bufio.NewScanner(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
		b   = strings.Builder{}

		t int
	)

	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {

		}
	}(out)

	in.Scan()
	_, err := fmt.Sscanf(in.Text(), "%d", &t)
	if err != nil {
		return
	}

	for tc := 0; tc < t; tc++ {
		var (
			n             int
			oddWordChars  = make(map[string]int)
			evenWordChars = make(map[string]int)
			wordKeys      = make(map[struct{ odd, even string }]int)
		)

		in.Scan()

		_, err = fmt.Sscanf(in.Text(), "%d", &n)
		if err != nil {
			return
		}

		for i := 0; i < n; i++ {
			in.Scan()

			s := in.Text()

			keyOdd := getKey(s, 0)
			keyEven := getKey(s, 1)

			oddWordChars[keyOdd]++
			evenWordChars[keyEven]++

			keys := struct{ odd, even string }{keyOdd, keyEven}
			wordKeys[keys]++
		}

		oddTotalSum := wordsCalc(oddWordChars)
		evenTotalSum := wordsCalc(evenWordChars)

		sumBoth := 0
		for key, m := range wordKeys {
			if len(key.odd) > 0 && len(key.even) > 0 {
				sumBoth += m * (m - 1) / 2
			}
		}

		total := oddTotalSum + evenTotalSum - sumBoth

		b.WriteString(strconv.Itoa(total))
		b.WriteString("\n")

		_, err = out.WriteString(b.String())
		if err != nil {
			return
		}

		b.Reset()
	}

}
