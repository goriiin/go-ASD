package main

func DigPow(n, p int) int {

	copyN := n
	var newN int

	for n > 0 {
		newN = newN*10 + n%10
		n /= 10
	}

	sum := 0
	for newN > 0 {
		sum += pow(newN%10, p)

		p++
		newN /= 10
	}

	if (sum/copyN)*copyN == sum {
		return sum / copyN
	}

	return -1
}

func pow(n, k int) int {
	if k < 0 {
		return 0
	} else if k == 0 {
		return 1
	} else if k%2 == 0 {
		half := pow(n, k/2)
		return half * half
	} else {
		return pow(n, k-1) * n
	}
}
