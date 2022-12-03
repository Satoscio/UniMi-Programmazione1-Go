/*
    Scrivere un programma somme.go che genera 10 numeri interi casuali tra
    0 e 10, li stampa, e stampa la somma dei valori generati.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n, somma int

    for i := 0; i < MAX; i++ {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        somma += n
    }
    fmt.Println()
    fmt.Println("Somma dei numeri generati:", somma)
}
