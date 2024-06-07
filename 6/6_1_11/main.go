package main

import "fmt"

func isPalindome(text string) string {
	if len(text) == 1 {
		return "YES"
	}

	leftCursor := 0
	rightCursor := len(text) - 1
	for leftCursor <= rightCursor {
		if text[leftCursor] != text[rightCursor] {
			return "NO"
		}
		leftCursor++
		rightCursor--
	}
	return "YES"
}

func main() {
	var text string

	_, _ = fmt.Scan(&text)

	fmt.Println(isPalindome(text))
}
