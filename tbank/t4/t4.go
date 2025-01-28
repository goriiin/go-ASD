package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

type addToDiv struct {
	toA1     uint64
	toA2     uint64
	toA3     uint64
	toA1A2   uint64
	toA1A3   uint64
	toA2A3   uint64
	toA1A2A3 uint64
}

func scd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func scc(n1, n2 uint64) (uint64, bool) {
	gcd := scd(n1, n2)
	if gcd == 0 {
		return 0, false
	}

	if n1/gcd > math.MaxUint64/n2 {
		return 0, false
	}

	return (n1 / gcd) * n2, true
}

func getInfoToAdd(a, num uint64) uint64 {
	return a - num%a
}

func newAddToDiv(num, a1, a2, a3 uint64) (addToDiv, error) {
	sccA1A2, ok := scc(a1, a2)
	if !ok {
		return addToDiv{}, errors.New("a1 and a2 does not exist")
	}

	sccA1A3, ok := scc(a1, a3)
	if !ok {
		return addToDiv{}, errors.New("a1 and a3 does not exist")
	}

	sccA2A3, ok := scc(a2, a3)
	if !ok {
		return addToDiv{}, errors.New("a2 and a3 does not exist")
	}

	sccA1A2A3, ok := scc(sccA1A2, a3)
	if !ok {
		return addToDiv{}, errors.New("a1, a2 and a3 does not exist")
	}

	return addToDiv{
		toA1:     getInfoToAdd(a1, num),
		toA2:     getInfoToAdd(a2, num),
		toA3:     getInfoToAdd(a3, num),
		toA1A2:   sccA1A2 - num,
		toA1A3:   sccA1A3 - num,
		toA2A3:   sccA2A3 - num,
		toA1A2A3: sccA1A2A3 - num,
	}, nil
}

func scanUint(in *bufio.Scanner) (ans uint64) {
	in.Scan()

	_, err := fmt.Sscanf(in.Text(), "%d", &ans)
	if err != nil {
		log.Fatal(err)
	}

	return ans
}

func findMin(nums ...uint64) uint64 {
	res := nums[0]

	for _, n := range nums {
		if n < res {
			res = n
		}
	}

	return res
}

func main() {
	var (
		minAdd = addToDiv{
			toA1:     math.MaxUint64,
			toA2:     math.MaxUint64,
			toA3:     math.MaxUint64,
			toA1A2:   math.MaxUint64,
			toA1A3:   math.MaxUint64,
			toA2A3:   math.MaxUint64,
			toA1A2A3: math.MaxUint64,
		}

		in  = bufio.NewScanner(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
	)

	in.Split(bufio.ScanWords)

	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	n, x, y, z := scanUint(in), scanUint(in), scanUint(in), scanUint(in)

	for i := 0; i < int(n); i++ {
		num := scanUint(in)

		add, err := newAddToDiv(num, x, y, z)
		if err != nil {
			log.Fatal(err)
		}

		if add.toA1 < minAdd.toA1 {
			minAdd.toA1 = add.toA1
		}

		if add.toA2 < minAdd.toA2 {
			minAdd.toA2 = add.toA2
		}

		if add.toA3 < minAdd.toA3 {
			minAdd.toA3 = add.toA3
		}

		if add.toA1A2 < minAdd.toA1A2 {
			minAdd.toA1A2 = add.toA1A2
		}

		if add.toA1A3 < minAdd.toA1A3 {
			minAdd.toA1A3 = add.toA1A3
		}

		if add.toA2A3 < minAdd.toA2A3 {
			minAdd.toA2A3 = add.toA2A3
		}

		if add.toA1A2A3 < minAdd.toA1A2A3 {
			minAdd.toA1A2A3 = add.toA1A2A3
		}
	}

	res := findMin(
		minAdd.toA1A2A3,
		minAdd.toA1+minAdd.toA2+minAdd.toA1,
		minAdd.toA1A2+minAdd.toA3,
		minAdd.toA1A3+minAdd.toA2,
		minAdd.toA2A3+minAdd.toA1,
	)

	_, err := fmt.Fprintln(out, res)
	if err != nil {
		return
	}
}
