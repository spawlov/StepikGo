package main

import "fmt"

func main() {
	var x, y int
	var operate string
	fmt.Scan(&x, &y, &operate)
	if operate == "+" {
		fmt.Println(x + y)
	} else if operate == "-" {
		fmt.Println(x - y)
	} else if operate == "*" {
		fmt.Println(x * y)
	} else if operate == "/" {
		if y == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println(float64(x) / float64(y))
		}
	} else {
		fmt.Println("Неверная операция")
	}
}
