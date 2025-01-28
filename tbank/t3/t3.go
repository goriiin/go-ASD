package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// префиксные суммы

func scanInt(in *bufio.Scanner) (ans int, err error) {
	in.Scan()

	_, err = fmt.Sscanf(in.Text(), "%d", &ans)

	return ans, err
}

func main() {
	var (
		in  = bufio.NewScanner(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
	)

	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	in.Split(bufio.ScanWords)

	n, err := scanInt(in)
	if err != nil {
		log.Fatal(err)
	}

	m, err := scanInt(in)
	if err != nil {
		log.Fatal(err)
	}

	days := make([]int, 0, n)
	numMap := make(map[int]int, n)

	var num int

	for i := 0; i < n; i++ {
		num, err = scanInt(in)

		_, ok := numMap[num]
		if ok {
			numMap[num]++

			continue
		}

		days = append(days, num)
		numMap[num]++
	}

	first, second := days[0], days[1]
	numMap[first]--
	numMap[second]--

	sort.Ints(days)

	prefSum := make([]int, len(days)+1)

	for i := 0; i < len(days); i++ {
		prefSum[i+1] = prefSum[i] + numMap[days[i]]
	}

	var (
		indF = sort.SearchInts(days, first)
		indS = sort.SearchInts(days, second)

		counter = prefSum[indS+1] - prefSum[indF]
	)

	answer := 0

}
