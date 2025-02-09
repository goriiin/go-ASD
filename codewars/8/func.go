package main

func ArrayDiff(a, b []int) []int {
	m := make(map[int]struct{}, len(b))

	for _, num := range b {
		m[num] = struct{}{}
	}

	res := make([]int, 0, len(a))

	for _, num := range a {
		_, ok := m[num]
		if !ok {
			res = append(res, num)
		}
	}

	return res
}
