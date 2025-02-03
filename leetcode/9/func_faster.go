package main

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	copyX := x
	reversed := 0
	for copyX > 0 {
		reversed = reversed*10 + copyX%10

		copyX /= 10
	}

	return reversed == x
}
