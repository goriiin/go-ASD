package main

import "fmt"

func main() {
	var (
		buttons = make([]bool, 10)
		a, b, c int32
		input   string
	)

	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		return
	}
	buttons[a], buttons[b], buttons[c] = true, true, true

	_, err = fmt.Scan(&input)
	if err != nil {
		return
	}

	for _, ch := range input {
		if !buttons[ch-'0'] {
			buttons[ch-'0'] = true
		}
	}

	counter := 0
	for _, flag := range buttons {
		if flag {
			counter++
		}
	}

	fmt.Println(counter - 3)
}
