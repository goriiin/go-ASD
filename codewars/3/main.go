package main

import "fmt"

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2)) // "abigailtheta"

	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1)) //  "oocccffuucccjjjkkkjyyyeehh"

	fmt.Println(LongestConsec([]string{}, 3)) // ""

	fmt.Println(LongestConsec([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2)) // "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck"
}
