package main

import (
	"fmt"
	"unicode"
)

func main() {
	var char rune

	_, _ = fmt.Scanf("%c", &char)

	if unicode.IsUpper(char) {
		fmt.Println(string(unicode.ToLower(char)))
	} else {
		fmt.Println(string(unicode.ToUpper(char)))
	}
}
