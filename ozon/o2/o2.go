package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func zeroCount(n int) int {
	if n == 0 {
		return 1
	}

	count := 1
	pow := 10
	temp := 1

	for pow <= n {
		count = temp
		pow *= 10
		temp++
	}

	return count
}

func findMax(d int, s string, n int) (int, error) {
	maxLen := len(s)
	result := 0

	for k := maxLen; k >= 1; k-- {
		candidate := strings.Repeat(strconv.Itoa(d), k)

		if len(candidate) < maxLen {
			num, err := strconv.Atoi(candidate)
			if err != nil {
				return 0, err
			}

			if num <= n {
				result = k

				break
			}
		} else {
			if candidate <= s {
				result = k

				break
			}
		}
	}

	return result, nil
}

func build(d, k, maxLen int) string {
	candidate := strings.Repeat(strconv.Itoa(d), k)

	if len(candidate) < maxLen {
		candidate += strings.Repeat("9", maxLen-k)
	} else {
		candidate = candidate[:k] + strings.Repeat("9", maxLen-k)
	}

	return candidate
}

func occurrencesDegits(candidate string, d int) int {
	counter := 0
	for _, c := range candidate {
		if c == rune('0'+d) {
			counter++
		}
	}

	return counter
}

func countMaxDigit(d, n int) (int, error) {
	if n < d {
		return 0, nil
	}

	if d == 0 {
		return zeroCount(n), nil
	}

	s := strconv.Itoa(n)
	maxLen := len(s)

	result, err := findMax(d, s, n)
	if err != nil {
		return 0, err
	}

	for k := maxLen; k >= 1; k-- {
		candidate := build(d, k, maxLen)
		if candidate > s {
			continue
		}

		occurrences := occurrencesDegits(candidate, d)
		if occurrences > result {
			result = occurrences
		}

		break
	}

	return result, nil
}

func main() {
	var (
		in  = bufio.NewScanner(os.Stdin)
		out = bufio.NewWriter(os.Stdout)

		n, sum int
	)

	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			panic(err)
		}
	}(out)

	in.Scan()

	t, err := strconv.Atoi(in.Text())
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		in.Scan()

		n, err = strconv.Atoi(in.Text())
		if err != nil {
			return
		}

		total := 0

		for d := 0; d <= 9; d++ {
			sum, err = countMaxDigit(d, n)
			if err != nil {
				return
			}

			total += sum
		}

		_, err = fmt.Fprintln(out, total)
		if err != nil {
			return
		}
	}
}
