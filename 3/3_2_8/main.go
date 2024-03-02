package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if -3 >= x || x >= 7 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
