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
	x, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	z, _ := strconv.Atoi(scanner.Text())

	fmt.Println(x+y+z, x*y*z)
}
