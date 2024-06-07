package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		k int; 
		text string;
		result []rune
	)

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_ = scanner.Scan()
	text = scanner.Text()
	text = strings.TrimSpace(text)

	_ = scanner.Scan()
	k, _ = strconv.Atoi(scanner.Text())


	runes := []rune(text)
	for i, symbol := range runes {
		if i != k-1 {
			result = append(result, symbol)
		}
	}
	fmt.Fprintln(writer, string(result))
}
