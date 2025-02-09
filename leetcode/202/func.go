package main

func isHappy(n int) bool {
	if n == 1 || n == 7 {
		return true
	} else if n < 10 {
		return false
	}

	prev := -1
	var curr int

	for {
		curr = 0

		for n > 0 {
			a := n % 10
			curr += a * a
			n /= 10
		}

		if curr == 1 || curr == 7 {
			return true
		}

		if curr == prev || curr < 10 {
			break
		}

		prev = curr
		n = curr
	}

	return false
}
