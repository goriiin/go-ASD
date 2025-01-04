package main

import "fmt"

func main() {
	var A, B, C uint64

	_, err := fmt.Scan(&A, &B, &C)
	if err != nil {
		return
	}

	if A >= B+C || B >= A+C || C >= A+B {
		fmt.Println("NO")

		return
	}

	fmt.Println("YES")
}
