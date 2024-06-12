package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string

	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text = strings.TrimSpace(scanner.Text())
	runes := []rune(text) 
	
	for idx, char := range runes {
	    if char == rune('e') {
	        runes[idx] = rune('i')
	    }
	}
	fmt.Println(string(runes))
}
