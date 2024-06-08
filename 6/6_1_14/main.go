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
	text = scanner.Text()
	text = strings.TrimSpace(text)

	fmt.Println(strings.Join(strings.Fields(text), " "))
}
