package main

import "fmt"

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	mas := []byte(isbn)
	total := 0

	for i, ch := range mas {
		fmt.Print(string(ch), "*", i+1, "+")
		if ch >= '0' && ch <= '9' {
			total += int(ch-'0') * (i + 1)
		} else {
			total += 10 * (i + 1)
		}

	}
	fmt.Println(total)

	return total%11 == 0
}
