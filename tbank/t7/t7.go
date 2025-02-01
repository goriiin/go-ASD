package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanInt(in *bufio.Scanner) (ans int) {
	in.Scan()

	_, err := fmt.Sscanf(in.Text(), "%d", &ans)
	if err != nil {
		log.Fatal(err)
	}

	return ans
}

func main() {
	const MOD = 998244353

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

	n, k := scanInt(in), scanInt(in)
	mas := make([]int, n)
	s := 2
	for i := 0; i < n; i++ {
		mas[i] = scanInt(in)
		s++
	}

	masX := make([]int, k+1)

	masX[0] = n % MOD
	for _, value := range mas {
		p := 1

		for t := 1; t <= k; t++ {
			p = (p * value) % MOD
			masX[t] = (masX[t] + p) % MOD
		}
	}

	C := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		C[i] = make([]int, k+1)
	}

	C[0][0] = 1
	for r := 1; r <= k; r++ {
		C[r][0] = 1
		for m := 1; m <= r; m++ {
			C[r][m] = (C[r-1][m-1] + C[r-1][m]) % MOD
		}
	}

	twoPows := make([]int, k+1)
	twoPows[0] = 1
	for r := 1; r <= k; r++ {
		twoPows[r] = (twoPows[r-1] * 2) % MOD
	}

	inv2 := (MOD + 1) / 2

	for r := 1; r <= k; r++ {
		S := 0

		for m := 0; m <= r; m++ {
			term := (((C[r][m] * masX[m]) % MOD) * masX[r-m]) % MOD
			S = (S + term) % MOD
		}

		val := (((S*inv2)%MOD-((twoPows[r-1]*masX[r])%MOD))%MOD + MOD) % MOD

		_, err := out.WriteString(fmt.Sprintf("%d\n", val))
		if err != nil {
			return
		}
	}
}
