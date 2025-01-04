package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a, b, c float64
		D       float64
	)

	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		return
	}

	D = b*b - 4*a*c
	if D < 0 {
		fmt.Println(0)

		return
	} else if D == 0 {
		fmt.Println(1)

		fmt.Println((-b) / (2 * a))
	} else {
		fmt.Println(2)

		sqrtD := math.Sqrt(D)

		x1 := (-b + sqrtD) / (2 * a)
		x2 := (-b - sqrtD) / (2 * a)
		if x1 < x2 {
			fmt.Println(x1, x2)

			return
		}

		fmt.Println(x2, x1)
	}
}
