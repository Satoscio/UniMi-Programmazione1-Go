package main

import "fmt"

func main() {
    var x string
    fmt.Scan(&x)
    fmt.Println(x)
    for i, r := range x {
        fmt.Println(i, "\t", r, "\t", string(r))
    }
}
