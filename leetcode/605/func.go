package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	var counter int

	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		counter++

		return counter >= n
	}

	if len(flowerbed) > 1 && !(flowerbed[0] == 1 || flowerbed[1] == 1) {
		flowerbed[0] = 1
		counter++
	}

	for i := 1; i < len(flowerbed); i++ {
		if i == len(flowerbed)-1 && !(flowerbed[i] == 1 || flowerbed[i-1] == 1) {
			flowerbed[i] = 1
			counter++
		} else if !(flowerbed[i] == 1 || flowerbed[i-1] == 1 || flowerbed[i+1] == 1) {
			flowerbed[i] = 1
			counter++
		}
	}

	return counter >= n
}
