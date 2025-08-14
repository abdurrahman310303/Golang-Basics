package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Abdur Rahman"
	vowels := "aeiou"
	word = strings.ToLower(word)
	count := 0

	for _, letter := range word {
		if strings.ContainsRune(vowels, letter) {
			count++
		}
	}

	fmt.Println(count)
}
