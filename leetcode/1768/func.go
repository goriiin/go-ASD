package main

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) == 0 {
		return word2
	}

	if len(word2) == 0 {
		return word1
	}

	mas := make([]byte, 0, len(word1)+len(word2))
	var i, j int

	mas1, mas2 := []byte(word1), []byte(word2)

	for {
		if !(i < len(word1)) && !(j < len(word2)) {
			break
		} else if i < len(word1) && j < len(word2) {
			mas = append(mas, mas1[i])
			mas = append(mas, mas2[j])

			i++
			j++
		} else if i < len(word1) {
			mas = append(mas, mas1[i])

			i++
		} else {

			mas = append(mas, mas2[j])

			j++
		}
	}

	return string(mas)
}
