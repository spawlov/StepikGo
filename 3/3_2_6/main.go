package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if 99 < x && x < 1000 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
