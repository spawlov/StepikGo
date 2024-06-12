package main

import "fmt"

func main() {
	var code string

	_, _ = fmt.Scan(&code)

	result := []rune{}
	for i := 0; i < len(code); i++ {
		if code[i] == '.' {
			result = append(result, '0')
			continue
		} else if code[i] == '-' && code[i+1] == '.' {
			result = append(result, '1')
			i++
			continue
		} else {
			result = append(result, '2')
			i++
		}
	}
	fmt.Println(string(result))
}
