package main

import (
	"strings"
)

func customSortString(S string, T string) string {
	var result string

	for i := 0; i < len(S); i++ {
		ch := ""
		for j := 0; j < len(T); j++ {
			if S[i] == T[j] {
				ch = string(T[j])
				result += ch
			}
		}
		T = strings.Replace(T, ch, "", -1)
	}
	result += T
	return result
}

func main() {
	customSortString("kqep", "pekeq")
}
