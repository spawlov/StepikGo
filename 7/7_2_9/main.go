package main

import "fmt"

func main() {
	var wordOne, wordTwo string

	_, _ = fmt.Scan(&wordOne, &wordTwo)

	fmt.Println(countChar(wordOne, "b") + countChar(wordTwo, "b"))
}

func countChar(word, symbol string) int {
	sum := 0
	slice := []rune(word)
	for _, char := range slice {
		if string(char) == symbol {
			sum++
		}
	}
	return sum
}
