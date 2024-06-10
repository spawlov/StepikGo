package main

import (
	"fmt"
)

func main() {
	var char rune

	_, _ = fmt.Scanf("%c", &char)

	for i := 97; i < int(char); i++ {
		fmt.Print(string(int32(i)), " ")
	}
	fmt.Println(string(char))
}
