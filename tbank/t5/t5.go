package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	prefix := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + a[i-1]
	}

	total := int64(n * (n + 1) / 2)
	sum := int64(0)

	for i := 1; i < n; i++ {
		A := prefix[i] - s
		B := prefix[i+1] - s

		left := sort.Search(i, func(k int) bool {
			return prefix[k] >= A
		})

		right := sort.Search(i, func(k int) bool {
			return prefix[k] >= B
		})

		cnt := right - left
		if cnt > 0 {
			sum += int64(cnt) * int64(n-i)
		}
	}

	fmt.Println(sum + total)
}
