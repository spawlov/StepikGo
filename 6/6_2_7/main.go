package main

import (
	"fmt"
)

func main() {
	var char rune

	_, _ = fmt.Scanf("%c", &char)

	for i := int(char); i <= int(rune('z')); i++ {
		fmt.Print(string(int32(i)), " ")
	}
	fmt.Print("\n")
}
