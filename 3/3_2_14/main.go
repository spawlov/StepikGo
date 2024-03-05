package main

import (
    "fmt"
    "math"
)

func main() {
    var x1, x2, y1, y2 int
    fmt.Scan(&x1, &y1, &x2, &y2)
    if math.Abs(float64(x2-x1)) == math.Abs(float64(y2-y1)) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
