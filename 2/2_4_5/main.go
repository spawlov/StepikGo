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
	num1, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	num2, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	num3, _ := strconv.Atoi(scanner.Text())

	fmt.Println(num1 * num2 * num3)
}
