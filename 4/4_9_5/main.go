package main

import "fmt"

func main() {
        count := 0

        for i := 100; i <= 999; i++ {
                digit1 := i / 100       // Первая цифра числа
                digit2 := (i / 10) % 10 // Вторая цифра числа
                digit3 := i % 10        // Третья цифра числа

                if digit1+digit2+digit3 == 8 {
                        count++
                }
        }

        fmt.Println("Количество чисел с суммой цифр равной 8: ", count)
}