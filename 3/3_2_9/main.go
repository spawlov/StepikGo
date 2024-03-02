package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	one := x % 10
	two := x / 10 % 10
	three := x / 100 % 10
	if one != two && one != three && two != three {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
