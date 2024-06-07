package main

import "fmt"

func main() {
	var n, height int

	_, _ = fmt.Scan(&n)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&heights[i])
	}

	_, _ = fmt.Scan(&height)

	position := n
	for no, heightId := range heights {
		if heightId < height {
			position = no
			break
		}
	}
	fmt.Println(position + 1)
}
