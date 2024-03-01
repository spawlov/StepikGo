package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	num1, _ := strconv.ParseFloat(scanner.Text(), 32)
	_ = scanner.Scan()
	num2, _ := strconv.ParseFloat(scanner.Text(), 32)
	fmt.Println((num1 + num2) / 2)
}
