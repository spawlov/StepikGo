package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	x1, _ := strconv.ParseFloat(scanner.Text(), 32)
	_ = scanner.Scan()
	y1, _ := strconv.ParseFloat(scanner.Text(), 32)
	_ = scanner.Scan()
	x2, _ := strconv.ParseFloat(scanner.Text(), 32)
	_ = scanner.Scan()
	y2, _ := strconv.ParseFloat(scanner.Text(), 32)

	fmt.Println(math.Abs(x1-x2) + math.Abs(y1-y2))
}
