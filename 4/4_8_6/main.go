package main

import . "fmt"

func main() {
	var n, c, d int
	_, _ = Scan(&n, &c, &d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			Println(i)
			break
		}
	}
}