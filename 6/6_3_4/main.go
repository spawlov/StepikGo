package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var firstStr, secondStr string

	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	firstStr = strings.TrimSpace(scanner.Text())
	_ = scanner.Scan()
	secondStr = strings.TrimSpace(scanner.Text())

	if firstStr[0] == secondStr[len(secondStr)-1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
