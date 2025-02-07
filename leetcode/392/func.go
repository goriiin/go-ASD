package main

import "fmt"

func isSubsequence(s string, t string) bool {
	chMap := make(map[byte][]int, 30)

	tBytes := []byte(t)
	for i, b := range tBytes {
		chMap[b] = append(chMap[b], i)
	}

	sBytes := []byte(s)
	var iterator int
	var ok bool
	for _, b := range sBytes {
		mas := chMap[b]
		iterator, ok = binSearch(mas, iterator)
		if !ok {
			return false
		}
		iterator++
	}

	return true
}

func binSearch(mas []int, target int) (int, bool) {
	l, r := 0, len(mas)-1
	result := -1

	for l <= r {
		mid := (l + r) / 2

		//fmt.Println("for: ", mid, mas[mid])
		if mas[mid] >= target {
			result = mas[mid]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println("input:", mas, target)

	fmt.Println("result:", result)

	if result != -1 {
		return result, true
	}

	return 0, false
}
