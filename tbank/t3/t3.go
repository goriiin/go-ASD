package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func scanInt(in *bufio.Scanner) (ans int64) {
	in.Scan()

	_, err := fmt.Sscanf(in.Text(), "%d", &ans)
	if err != nil {
		log.Fatal(err)
	}

	return ans
}

func main() {
	var (
		in  = bufio.NewScanner(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
	)

	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			log.Fatalln(err)
		}
	}(out)

	in.Split(bufio.ScanWords)

	var (
		n, m = scanInt(in), scanInt(in)
		a    = make([]int64, n)
	)

	for i := 0; i < int(n); i++ {
		a[i] = scanInt(in)
	}

	a1, a2 := a[0], a[1]

	days := a[2:]
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	var minCost int64 = math.MaxInt64

	for i := 0; i <= len(days)-int(m); i++ {
		j := i + int(m) - 1
		costLeft := a1 - days[i]
		if costLeft < 0 {
			costLeft = 0
		}

		costRight := days[j] - a2
		if costRight < 0 {
			costRight = 0
		}

		total := costLeft + costRight
		if total < minCost {
			minCost = total
		}
	}

	_, err := out.WriteString(fmt.Sprintf("%d\n", minCost))
	if err != nil {
		return
	}
}
