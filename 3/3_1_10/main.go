package main

import "fmt"

func main() {
	var passw1, passw2 string
	fmt.Scan(&passw1, &passw2)
	if passw1 == passw2 {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}
