package main

import "fmt"

func main() {
	var number float32
	const pi = 3.14159

	fmt.Scan(&number)
	fmt.Println(number * number * pi)
}
