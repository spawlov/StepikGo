package main

import (
	"fmt"
	"regexp"
)

func isId(text string) string {
	result := "NO"
	regex := regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
	if regex.MatchString(text) {
		result = "YES"
	}
	return result
}

func main() {
	var text string

	_, _ = fmt.Scan(&text)

	fmt.Println(isId(text))
}
