/*
    Scrivere un programma massimo.go che genera 10 numeri interi casuali tra
    10 e 30, li stampa, e stampa il massimo valore generato.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    const NUMERI, MIN, MAX = 10, 10, 30
    rand.Seed(time.Now().Unix())
    var n int
    max := 0

    for i := 0; i < NUMERI; i++ {
        n = rand.Intn(MAX - MIN) + MIN
        if n > max {
            max = n
        }
        fmt.Print(n, " ")
    }
    fmt.Println()
    fmt.Println("Il massimo è:", max)
}
