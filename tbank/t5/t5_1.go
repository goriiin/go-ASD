package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
			log.Fatal(err)
		}
	}(out)

	in.Split(bufio.ScanWords)
	n, s := scanInt(in), scanInt(in)
	masA := make([]int64, n+1)

	for i := 1; i <= int(n); i++ {
		masA[i] = scanInt(in)
	}

	prefSumMas := make([]int64, n+1)
	for i := 1; i <= int(n); i++ {
		prefSumMas[i] = prefSumMas[i-1] + masA[i]
	}

	dp := make([]int64, n+1)

	start := 1
	var result int64

	for i := 1; i <= int(n); i++ {
		for prefSumMas[i]-prefSumMas[start-1] > s {
			start++
		}

		pI := start

		dp[i] = dp[pI-1] + int64(i)
		result += dp[i]
	}

	_, err := out.WriteString(fmt.Sprintf("%d\n", result))
	if err != nil {
		return
	}
}
