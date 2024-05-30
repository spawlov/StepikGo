package main

import "fmt"

func main() {
	var a, b, result int

	_, _ = fmt.Scan(&a, &b)

	found := false

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			result = i
			found = true
			break
		}
	}
	if found {
		fmt.Println(result)
	} else {
		fmt.Println("NO")
	}
}
