package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	mas := make([]int8, 0, 20)

	for x > 0 {
		mas = append(mas, int8(x%10))

		x /= 10
	}

	var (
		l      = 0
		r      = len(mas) - 1
		length = len(mas)
	)

	for l != length/2 && (r >= 0 || l < length-1) {
		if mas[l] != mas[r] {
			return false
		}

		l++
		r--
	}

	return true
}
