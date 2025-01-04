package main

import (
	"fmt"
)

func main() {
	var A, B uint64

	_, err := fmt.Scan(&A, &B)
	if err != nil {
		return
	}

	fmt.Println(A + B)
}
