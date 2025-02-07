package main

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	} else if n%2 != 0 || n <= 0 {
		return false
	}

	mas := make([]int, 31)
	mas[0] = 1

	for i := 1; i < len(mas); i++ {
		mas[i] = mas[i-1] * 2
	}

	return binSearch(mas, n)
}

func binSearch(mas []int, target int) bool {
	l, r := 0, len(mas)-1

	for l <= r {
		mid := l + (r-l)/2

		if mas[mid] == target {
			return true
		} else if mas[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
