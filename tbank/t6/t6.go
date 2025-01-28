package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return 1
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	if n < 3 {
		fmt.Println(0)
		return
	}

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		points[i] = Point{x, y}
	}

	maxM := 0
	for i := 0; i < n; i++ {
		slopeCounts := make(map[string]int)
		xi, yi := points[i].x, points[i].y
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			xj, yj := points[j].x, points[j].y
			dx := xj - xi
			dy := yj - yi

			g := gcd(int(math.Abs(float64(dx))), int(math.Abs(float64(dy))))
			if g == 0 {
				g = 1
			}
			dxReduced := dx / g
			dyReduced := dy / g

			if dxReduced < 0 || (dxReduced == 0 && dyReduced < 0) {
				dxReduced = -dxReduced
				dyReduced = -dyReduced
			}

			key := fmt.Sprintf("%d,%d", dxReduced, dyReduced)
			slopeCounts[key]++
		}

		currentMax := 0
		for _, v := range slopeCounts {
			if v > currentMax {
				currentMax = v
			}
		}
		currentMax += 1 // +1 для текущей точки i
		if currentMax > maxM {
			maxM = currentMax
		}
	}

	m := maxM
	if m < 3 {
		fmt.Println(n / 3)
	} else {
		k := n - m
		t1 := min(m/2, k)
		mPrime := m - 2*t1
		kPrime := k - t1
		remaining := mPrime + kPrime

		t2 := 0
		if remaining >= 3 && kPrime > 0 {
			t2 = remaining / 3
		}
		total := t1 + t2
		fmt.Println(min(total, n/3))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
