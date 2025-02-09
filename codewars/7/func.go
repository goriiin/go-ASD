package main

func Multiple3And5(number int) int {
	if number < 3 {
		return 0
	}

	sum := 0
	i, j := 1, 1
	for {
		n3, n5 := i*3, j*5

		if n3 >= number && n5 >= number {
			break
		}

		if n3 < number {
			if i%5 != 0 {
				sum += n3
			}

			i++
		}

		if n5 < number {
			sum += n5

			j++
		}

	}

	return sum
}
