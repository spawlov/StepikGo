package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cost int
	scanner := bufio.NewScanner(os.Stdin)

	_ = scanner.Scan()
	rub, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	kop, _ := strconv.Atoi(scanner.Text())
	_ = scanner.Scan()
	pie, _ := strconv.Atoi(scanner.Text())

	cost = (rub*100 + kop) * pie

	fmt.Println(cost/100, cost%100)
}
