package main

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		swap(s, i, len(s)-i-1)
	}
}

func swap(s []byte, l, r int) {
	t := s[l]

	s[l], s[r] = s[r], t
}
