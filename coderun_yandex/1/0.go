package main

import (
	"fmt"
)

func avgV1(num1, num2, num3 int) int {
	if num1 > num2 {
		num1, num2 = num2, num1
	}

	if num3 < num2 {
		num3, num2 = num2, num3
	}

	if num1 > num2 {
		num1, num2 = num2, num1
	}

	return num2
}

func main() {
	var num1, num2, num3 int

	_, err := fmt.Scan(&num1, &num2, &num3)
	if err != nil {
		return
	}

	n := avg(num1, num2, num3)

	fmt.Println(n)
}
