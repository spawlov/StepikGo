package main

import "fmt"

func main() {
	var n, m string

	_, _ = fmt.Scan(&n, &m)

	fmt.Println(len(n) * len(m))
}
