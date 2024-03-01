package main

import "fmt"

func main() {
	x := 8 + 2*5
	y := (x % 10) + 14
	x = (y / 10) + 3
	c := x - y
	fmt.Println(c)
}
