package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	_, _ = fmt.Scan(&n)

	list := make([]string, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&list[i])
	}

	for _, text := range list {
		if len(text) > 10 {
			fmt.Println(string(text[0]) + strconv.Itoa(len(text)-2) + string(text[len(text)-1]))
		} else {
			fmt.Println(text)
		}
	}
}
