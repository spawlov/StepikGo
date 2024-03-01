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
	n, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println(k / n)
}
