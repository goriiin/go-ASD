package main

import "fmt"

func main() {
	var num1, num2, num3 int

	_, err := fmt.Scan(&num1, &num2, &num3)
	if err != nil {
		return
	}

	n := avgV1(num1, num2, num3)

	//n := avgV2(num1, num2, num3)

	fmt.Println(n)

}
