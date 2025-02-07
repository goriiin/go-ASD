package main

func plusOne(digits []int) []int {

	digits[len(digits)-1] += 1
	mem := digits[len(digits)-1] / 10
	digits[len(digits)-1] %= 10

	for i := len(digits) - 2; i >= 0; i-- {
		if mem == 0 {
			break
		}

		digits[i] += mem
	}
	if mem > 0 {
		digits = append([]int{1}, digits...)

	}

	return digits
}
