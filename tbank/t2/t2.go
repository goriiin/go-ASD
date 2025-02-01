package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		in  = bufio.NewScanner(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
		b   = strings.Builder{}

		n           int
		powersOfTwo = make([]uint64, 64)
		lib         = make(map[uint64]uint64, 100)
	)

	defer func(printer *bufio.Writer) {
		err := printer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	for i := 0; i < len(powersOfTwo); i++ {
		powersOfTwo[i] = 1 << uint(i)
	}

	getNum := func(length, bigInd, midInd, smallInd int) uint64 {
		ans := powersOfTwo[length-bigInd-1] +
			powersOfTwo[length-midInd-1] +
			powersOfTwo[length-smallInd-1]

		return ans
	}

	in.Scan()
	_, err := fmt.Sscanf(in.Text(), "%d", &n)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var a uint64

		in.Scan()
		_, err = fmt.Sscanf(in.Text(), "%d", &a)
		if err != nil {
			return
		}

		if a < 7 {
			b.WriteString(strconv.FormatInt(int64(-1), 10))

			b.WriteString("\n")

			_, err = out.WriteString(b.String())
			if err != nil {
				return
			}

			b.Reset()

			continue
		}

		ans, ok := lib[a]
		if !ok {
			var bigInd, midInd, smallInd, counter int
			var length = bits.Len64(a)
			for i := length - 1; i >= 0; i-- {
				if a&(1<<uint(i)) != 0 {
					switch counter {
					case 0:
						bigInd = bits.Len64(a) - 1 - i
					case 1:
						midInd = bits.Len64(a) - 1 - i
					case 2:
						smallInd = bits.Len64(a) - 1 - i
					}
					counter++
					if counter == 3 {
						break
					}
				}
			}

			if counter > 2 {
				ans = getNum(length, bigInd, midInd, smallInd)
			} else {
				if midInd == 0 || midInd >= length-2 {
					ans = getNum(length, bigInd+1, bigInd+2, bigInd+3)
				} else {
					ans = getNum(length, bigInd, midInd+1, midInd+2)
				}
			}

			b.WriteString(strconv.FormatUint(ans, 10))
			lib[a] = ans
		} else {
			b.WriteString(strconv.FormatUint(lib[a], 10))
		}

		b.WriteString("\n")

		_, err = out.WriteString(b.String())
		if err != nil {
			return
		}

		b.Reset()
	}
}
