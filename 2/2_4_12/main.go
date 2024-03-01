package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Scan(&number)
	fmt.Println("Следующее за числом", number, "число:", number+1)
	fmt.Println("Для числа", number, "предыдущее число:", number-1)
}
