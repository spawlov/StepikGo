package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		text string
	)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_ = scanner.Scan()
	text = scanner.Text()
	text = strings.TrimSpace(text)

	result := "-1"
	runes := []rune(text)
	for _, symbol := range runes {
		if symbol == rune('x') || symbol == rune('w') {
			result = string(symbol)
			break
		}
	}
	fmt.Fprintln(writer, result)
}
