package main

import (
	. "fmt"
)

func main() {
	var n, firstMax, secondMax int

	_, _ = Scan(&firstMax)
	_, _ = Scan(n)

	if n > firstMax {
		secondMax = firstMax
		firstMax = n
	}

	for {
		_, _ = Scan(&n)
		if n == 0 {
			break
		}

		if n > firstMax {
			secondMax = firstMax
			firstMax = n
		} else if n > secondMax {
			secondMax = n
		}
	}
	Println(secondMax)
}
