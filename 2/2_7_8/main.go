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
	katet1, _ := strconv.ParseFloat(scanner.Text(), 64)
	_ = scanner.Scan()
	katet2, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println(math.Sqrt(math.Pow(katet1, 2) + math.Pow(katet2, 2)))
}
