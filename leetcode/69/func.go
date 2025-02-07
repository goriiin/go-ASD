package main

func mySqrt(x int) int {
	l, r := 0, x

	var mid int
	for l <= r {
		mid = l + (r-l)/2

		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1

			if (mid+1)*(mid+1) > x {
				return mid
			}
		}
	}

	return mid
}
