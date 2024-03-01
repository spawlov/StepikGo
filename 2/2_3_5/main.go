package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var bookName string
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	bookName = scanner.Text()
	fmt.Println(bookName, "- лучшая книга!")
}
