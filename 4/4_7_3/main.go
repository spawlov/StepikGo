package main

import (
	. "fmt"
	. "strconv"
	. "strings"
)

func main() {
	var number int
	_, _ = Scan(&number)
	counter := 0
	for i := 0; i <= number; i++ {
		counter += Count(Itoa(i), "7")
	}
	Println(counter)
}
