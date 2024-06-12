package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string

	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text = strings.TrimSpace(scanner.Text())

	fmt.Println(countVowels(text))
}

func countVowels(text string) int {
	result := 0
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, char := range text {
		if contains(vowels, string(char)) {
			result++
		}
	}
	return result
}

func contains(slice []string, char string) bool {
	for _, item := range slice {
		if item == char {
			return true
		}
	}
	return false
}
