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
	katet1, _ := strconv.ParseFloat(scanner.Text(), 64)
	_ = scanner.Scan()
	katet2, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println(katet1 * katet2 * 0.5)
}
