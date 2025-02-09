package main

func FirstNonRepeating(str string) string {
	m := make(map[int32]int, 2*len(str))

	for _, v := range str {
		m[v]++
		if 'a' < v && v <= 'z' {
			m['A'+v-'a']++
		} else if 'A' < v && v <= 'Z' {
			m['a'+v-'A']++
		}
	}

	var res string

	for _, v := range str {
		if m[v] == 1 {
			res = string(v)

			break
		}
	}

	return res
}
