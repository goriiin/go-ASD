package main

func Tribonacci(signature [3]float64, n int) []float64 {
	mas := make([]float64, 0, n+1)
	mas = append(mas, signature[0], signature[1], signature[2])

	for i := 3; i <= n; i++ {
		mas = append(mas, mas[i-1]+mas[i-2]+mas[i-3])
	}

	return mas
}
