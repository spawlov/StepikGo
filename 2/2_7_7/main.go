package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(number%10 + (number%100)/10 + (number%1000)/100)
}
