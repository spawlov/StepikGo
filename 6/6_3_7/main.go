package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var text string

	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text = strings.TrimSpace(scanner.Text())

	fmt.Println(strings.Join(extractDigits(text), " "))
}

func extractDigits(text string) []string {
	result := []string{}
	for _, char := range text {
		if unicode.IsDigit(char) {
			result = append(result, string(char))
		}
	}
	return result
}
