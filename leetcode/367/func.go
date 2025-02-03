package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	_, x := binSearch(0, num/1, num)

	return x
}

func binSearch(l, r, target int) (int, bool) {
	var mid int
	var flag bool

	for l <= r {
		mid = (l + r) / 2
		x := mid * mid
		if x == target {
			flag = true

			break
		} else if x > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return mid, flag
}
