package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	var (
		b = strings.Builder{}
	)
	b.WriteString(str1)
	b.WriteString(str2)

	l := b.String()
	b.Reset()

	b.WriteString(str2)
	b.WriteString(str1)
	r := b.String()

	if l != r {
		return ""
	}

	g := gcd(len(str1), len(str2))

	return string([]byte(str1)[:g])
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
