package t7

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

const MOD = 998244353

// Считаем, что формат входных данных такой:
//    n k
//    a_1 a_2 ... a_n
// И мы должны вывести f(1), f(2), ..., f(k).

func main() {
	// ---------------------------------------------------------
	// 0) Приготовим функции-утилиты для чтения и модульных операций
	// ---------------------------------------------------------

	// Быстрое чтение целых чисел (не супер-оптимально, но компактно)
	reader := bufio.NewReader(os.Stdin)
	readInts := func() []int {
		line, _ := reader.ReadString('\n')
		fields := make([]int, 0, 8)
		start := 0
		for i := 0; i < len(line); i++ {
			if line[i] == ' ' || line[i] == '\n' {
				if start < i {
					num, _ := strconv.Atoi(line[start:i])
					fields = append(fields, num)
				}
				start = i + 1
			}
		}
		return fields
	}

	add := func(a, b int) int {
		s := a + b
		if s >= MOD {
			s -= MOD
		}
		return s
	}
	sub := func(a, b int) int {
		s := a - b
		if s < 0 {
			s += MOD
		}
		return s
	}
	mul := func(a, b int) int {
		return int(int64(a) * int64(b) % MOD)
	}

	// ---------------------------------------------------------
	// 1) Чтение входных данных
	// ---------------------------------------------------------
	line := readInts()
	n, k := line[0], line[1]

	// Читаем массив a (n штук)
	aVals := make([]int, n)
	idx := 0
	for idx < n {
		chunk := readInts()
		for _, v := range chunk {
			aVals[idx] = v
			idx++
			if idx == n {
				break
			}
		}
	}

	// ---------------------------------------------------------
	// 2) Начинаем «бенчмарк»: снимем первую «точку» по времени/памяти
	// ---------------------------------------------------------
	var memStart, memAfterS, memAfterBinom, memAfterCalc runtime.MemStats

	// Снимаем «стартовую» статистику по памяти
	runtime.ReadMemStats(&memStart)
	tStart := time.Now()

	// ---------------------------------------------------------
	// 3) Предварительно вычисляем все степени a_i^r и суммы S[r].
	//    S[r] = Σ (a_i^r) mod. Нужно r=1..k.
	//
	//    Можем расширить, чтобы S[0] = n, см. пояснения ниже.
	// ---------------------------------------------------------
	S := make([]int, k+1) // S[r], r=1..k (пока индекс = r)
	for _, x := range aVals {
		cur := 1
		for r := 1; r <= k; r++ {
			cur = mul(cur, x)
			S[r] = add(S[r], cur)
		}
	}

	// (Если нужно учесть S[0]=n, можно вставить это вручную, см. пример ниже.)
	// Но в самой формуле мы можем скорректировать при подсчёте.

	// Замерим время после вычисления S
	tAfterS := time.Now()
	runtime.ReadMemStats(&memAfterS)

	// ---------------------------------------------------------
	// 4) Строим таблицу биномиальных коэффициентов binom[p][m] (p<=k)
	//    и массив степеней 2^p (mod).
	// ---------------------------------------------------------
	binom := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		binom[i] = make([]int, i+1)
		binom[i][0] = 1
		binom[i][i] = 1
		for j := 1; j < i; j++ {
			binom[i][j] = add(binom[i-1][j-1], binom[i-1][j])
		}
	}

	pow2 := make([]int, k+1)
	pow2[0] = 1
	for i := 1; i <= k; i++ {
		pow2[i] = mul(pow2[i-1], 2)
	}

	// Инверсия числа 2 по модулю 998244353
	inv2 := (MOD + 1) / 2 // 499122177

	// Замерим время после построения бинома
	tAfterBinom := time.Now()
	runtime.ReadMemStats(&memAfterBinom)

	// ---------------------------------------------------------
	// 5) Вычисляем f(p) для p=1..k.
	//
	//    Формула:
	//      f(p) = (1/2) * [ Σ_{m=0..p} (binom[p][m]*S[m]*S[p-m]) - 2^p * S[p] ]  (mod)
	//
	//    Чтобы учесть S[0], нам нужно завести S[0]=n, но можно «добавить» его в S,
	//    тогда S[r] сдвинутся, или обойтись, учтя отдельно члены m=0 и m=p.
	//    Для наглядности здесь сделаем «добавку» вручную.
	//    S[0] = n (т.к. a_i^0 = 1, суммируется n раз).
	// ---------------------------------------------------------
	S0 := n % MOD
	results := make([]int, k+1) // results[p] = f(p)

	for pVal := 1; pVal <= k; pVal++ {
		// sumT = Σ_{m=0..p} binom[p][m]*S[m]*S[p-m]
		//  где мы знаем:
		//    S[0] = n
		//    S[m] (m>=1) — есть в массиве (S[m]).
		//
		// В S[...] у нас сейчас S[r] соответствует r=1..k,
		// т. е. "S[m]" здесь = S[m], но для m=0 придётся подставить вручную S0.
		//
		var sumT int
		for m := 0; m <= pVal; m++ {
			var sm, spm int
			if m == 0 {
				sm = S0
			} else {
				sm = S[m]
			}
			if (pVal - m) == 0 {
				spm = S0
			} else {
				spm = S[pVal-m]
			}
			part := mul(binom[pVal][m], mul(sm, spm))
			sumT = add(sumT, part)
		}

		// Вычитаем 2^pVal * S[pVal] (S[pVal] — это действительно S[pVal], так как у нас r=1..k)
		subPart := mul(pow2[pVal], S[pVal])
		fP := sub(sumT, subPart)

		// Умножаем на 1/2
		fP = mul(fP, inv2)
		// Нормализуем
		if fP < 0 {
			fP += MOD
		}
		results[pVal] = fP
	}

	// Замерим время после вычисления f(p)
	tAfterCalc := time.Now()
	runtime.ReadMemStats(&memAfterCalc)

	// ---------------------------------------------------------
	// 6) Выводим результаты и одновременно выводим, как «бенчмарк»
	// ---------------------------------------------------------
	// 6.1. Выведем f(1)...f(k)
	writer := bufio.NewWriter(os.Stdout)
	for pVal := 1; pVal <= k; pVal++ {
		writer.WriteString(strconv.Itoa(results[pVal]))
		if pVal < k {
			writer.WriteByte(' ')
		}
	}
	writer.WriteByte('\n')
	writer.Flush()

	// 6.2. Выведем статистику по времени
	tEnd := time.Now()
	fmt.Println("=== Time measurements ===")
	fmt.Printf("Total time: %v\n", tEnd.Sub(tStart))
	fmt.Printf("Time to build S: %v\n", tAfterS.Sub(tStart))
	fmt.Printf("Time to build binom: %v\n", tAfterBinom.Sub(tAfterS))
	fmt.Printf("Time to calculate f(p): %v\n", tAfterCalc.Sub(tAfterBinom))
	fmt.Printf("Time after everything (including output): %v\n", tEnd.Sub(tStart))

	// 6.3. Выведем статистику по памяти (Alloc в байтах)
	fmt.Println("=== Memory usage (Alloc) ===")
	fmt.Printf("Before building S:  %d bytes\n", memStart.Alloc)
	fmt.Printf("After building S:   %d bytes\n", memAfterS.Alloc)
	fmt.Printf("After building bin: %d bytes\n", memAfterBinom.Alloc)
	fmt.Printf("After calc f(p):    %d bytes\n", memAfterCalc.Alloc)
	fmt.Printf("Memory used by S-building:    %d bytes\n", int64(memAfterS.Alloc)-int64(memStart.Alloc))
	fmt.Printf("Memory used by bin-building:  %d bytes\n", int64(memAfterBinom.Alloc)-int64(memAfterS.Alloc))
	fmt.Printf("Memory used by calculation:   %d bytes\n", int64(memAfterCalc.Alloc)-int64(memAfterBinom.Alloc))
}
