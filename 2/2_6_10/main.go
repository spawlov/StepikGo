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
	a, _ := strconv.ParseFloat(scanner.Text(), 32)
	_ = scanner.Scan()
	b, _ := strconv.ParseFloat(scanner.Text(), 32)
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println(a + b + c)
}
