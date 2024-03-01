package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	_ = scanner.Scan()
	sep := scanner.Text()
	_ = scanner.Scan()
	line1 := scanner.Text()
	_ = scanner.Scan()
	line2 := scanner.Text()
	_ = scanner.Scan()
	line3 := scanner.Text()

	fmt.Println(line1 + sep + line2 + sep + line3)
}
