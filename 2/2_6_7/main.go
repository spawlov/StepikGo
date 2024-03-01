package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	fmt.Println(num - float64(int(num)))
}
