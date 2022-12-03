package main

import "fmt"

func main() {
    var x string
    fmt.Scan(&x)
    fmt.Println(x)
    for i := 0; i < len(x); i++ {
        fmt.Println(i, "\t", x[i])
    }
}
