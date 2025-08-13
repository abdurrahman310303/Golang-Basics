package main

import "fmt"

// In Go, a rune is an alias for int32 and represents a Unicode code point.
// When we convert a string to []rune, each character in the string is converted
// to its Unicode code point value, which allows us to properly handle
// multi-byte characters like emojis or non-ASCII text.
//
// For example:
// "hello" -> []rune{104, 101, 108, 108, 111}
func reverseString(s string) string {
	// Convert string to rune slice to handle Unicode characters correctly
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(reverseString("Hello, World!"))
}
