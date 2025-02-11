package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxNum, maxInd int

	for i, c := range candies {
		if c > maxNum {
			maxNum = c
			maxInd = i
		}
	}

	res := make([]bool, len(candies))
	res[maxInd] = true

	for i, c := range candies {
		if c+extraCandies >= maxNum {
			res[i] = true
		}
	}

	return res
}
