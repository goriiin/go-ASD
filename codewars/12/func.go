package main

func RGB(r, g, b int) string {

	return toHEX(r) + toHEX(g) + toHEX(b)
}

func toHEX(a int) string {
	if a <= 0 {
		return "00"
	} else if a > 255 {
		return "FF"
	}

	mas := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	r := mas[a%16]
	a /= 16

	l := mas[a%16]

	return l + r
}
