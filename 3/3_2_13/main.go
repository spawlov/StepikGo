package main

import "fmt"

func main() {
    var x1, x2, y1, y2 int
    fmt.Scan(&x1, &y1, &x2, &y2)
    if x1 == x2 || y1 == y2 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
