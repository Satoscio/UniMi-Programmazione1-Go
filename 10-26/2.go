package main

import "fmt"

func main() {
    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)

    for n > 1 {
        if n % 2 == 0 {
            n /= 2
            fmt.Println(n)
        } else {
            n = n * 3 + 1
            fmt.Println(n)
        }
    }
}
