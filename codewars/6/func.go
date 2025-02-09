package main

func TowerBuilder(nFloors int) []string {
	if nFloors < 1 {
		return []string{}
	}

	res := make([]string, 0, nFloors)
	strMas := make([]byte, 1+(nFloors-1)*2)

	for i := 0; i < len(strMas); i++ {
		strMas[i] = ' '
	}

	mid := (1 + (nFloors-1)*2) / 2
	strMas[mid] = '*'

	res = append(res, string(strMas))

	for i := 1; i < nFloors; i++ {

		strMas[mid+i] = '*'
		strMas[mid-i] = '*'

		res = append(res, string(strMas))
	}

	return res
}
