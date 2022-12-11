package main

import (
    "fmt"
    "os"
)

func main() {
    x := os.Args[1:]
    for _, k := range x {
        fmt.Println(k)
    }
}
