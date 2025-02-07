package main

import "fmt"

func main() {

	s := "123"

	for _, ch := range s {
		fmt.Println(string(ch), s)
	}
}
