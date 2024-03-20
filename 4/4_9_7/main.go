package main

import "fmt"

func isPerfectNum(n int) bool {
        sum := 0

        for i := 1; i < n; i++ {
                if n%i == 0 {
                        sum += i
                }
        }

        return sum == n
}

func main() {
        count := 0
        num := 2

        for count < 3 {
                if isPerfectNum(num) {
                        fmt.Println(num)
                        count++
                }
                num++
        }
}