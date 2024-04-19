package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    max := 0
    cnt := 0
    for i := 1; i <= n; i++ {
        var k int
        fmt.Scan(&k)
        if k > max {
            max = k
            cnt = 1
        } else {
            if k == max {
                cnt++
            }
        }
    }
    fmt.Println(cnt)
}