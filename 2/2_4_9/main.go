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
	smartphone, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	case_, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	charger, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	headphones, _ := strconv.Atoi(scanner.Text())

	fmt.Println((smartphone + case_ + charger + headphones) * 3)
}
