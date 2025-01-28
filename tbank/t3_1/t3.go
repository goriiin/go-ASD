package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	// Читаем входные данные.
	// Формат ожидается такой:
	//    n m
	//    a[1] a[2] ... a[n]
	in := bufio.NewReader(os.Stdin)

	var n, m int
	_, err := fmt.Fscan(in, &n, &m)
	if err != nil {
		log.Fatal(err)
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Fscan(in, &a[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	// "Первые два" дня:
	a1 := a[0]
	a2 := a[1]

	// Обычные дни a[2..n-1]
	if n <= 2 {
		// Крайний случай: обычных дней нет или < m, результат может быть 0 или невозможен.
		// Но обычно по условию n >= 3, m <= n-2.
		fmt.Println(0)
		return
	}

	b := a[2:] // длина = n - 2
	sort.Ints(b)

	// Если m == 0 (если такое допускается), платить 0.
	// Но обычно m >= 1.
	if m <= 0 {
		fmt.Println(0)
		return
	}

	// Если (n-2) < m, то невозможно "набрать" m обычных дней,
	// но предполагаем, что тесты такие не дают.
	if len(b) < m {
		// Либо вывести 0, либо "невозможно". Решайте по условию.
		fmt.Println(0)
		return
	}

	// Перебираем подотрезки длины m в b.
	// i идёт от 0 до len(b)-m (включительно).
	// j = i + m - 1.
	// cost = max(0, a1 - b[i]) + max(0, b[j] - a2).
	// Берём минимум.
	const INF = 1_000_000_000_000_000_000
	minCost := INF

	limit := len(b) - m
	for i := 0; i <= limit; i++ {
		j := i + m - 1
		costLeft := a1 - b[i]
		if costLeft < 0 {
			costLeft = 0
		}
		costRight := b[j] - a2
		if costRight < 0 {
			costRight = 0
		}
		total := costLeft + costRight

		if total < minCost {
			minCost = total
		}
	}

	fmt.Println(minCost)
}
