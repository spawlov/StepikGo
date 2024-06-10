package main

import (
	"fmt"
	"unicode"
)

func main() {
	var char rune

	_, _ = fmt.Scanf("%c", &char)

	if unicode.IsDigit(char) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
