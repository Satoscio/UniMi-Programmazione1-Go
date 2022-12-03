package main

import "fmt"

func main() {
    var x, y rune
    x = 8984
    y = '\U0001F640'
    fmt.Println(x, string(x))
    fmt.Println(y, string(y))
}
