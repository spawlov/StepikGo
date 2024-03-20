package main

import "fmt"

func main() {
        found := false

        for i := 10; i <= 99; i++ {
                digit1 := i / 10 // Первая цифра числа
                digit2 := i % 10 // Вторая цифра числа

                if i == 2*digit1*digit2 {
                        if found {
                                fmt.Print(",")
                        }
                        fmt.Print(i)
                        found = true
                }
        }

        fmt.Println()
}