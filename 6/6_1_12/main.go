package main

import "fmt"

func main() {
	var text string

	_, _ = fmt.Scan(&text)

	frequency := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		frequency[text[i]]++
		if frequency[text[i]] == 2 {
			fmt.Println(string(text[i]))
			break
		}
	}
}
