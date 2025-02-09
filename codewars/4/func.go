package main

func FindUniq(arr []float32) float32 {
	m := make(map[float32]int)

	for _, num := range arr {
		m[num]++
	}

	var uniq float32

	for num, count := range m {
		if count == 1 {
			uniq = num
		}
	}

	return uniq
}
