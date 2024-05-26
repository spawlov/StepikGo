package main

import (
	. "fmt"
)

func main() {
	var N int
	_, _ = Scan(&N)

	maxSpeed := 0
	fixSpeed := "NO"

	for i := 0; i < N; i++ {
		var speed int

		_, _ = Scan(&speed)
		if speed > maxSpeed {
			maxSpeed = speed
		}
		if speed < 30 {
			fixSpeed = "YES"
		}
	}
	Print(maxSpeed, " ", fixSpeed, "\n")
}
