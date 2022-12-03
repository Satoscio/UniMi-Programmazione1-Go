/*
    Scrivere un programma es10.go che genera numeri casuali tra 1 e 10,
    stampandoli, fino a quando la loro somma raggiunge un obiettivo fissato
    (TARGET), ad esempio 50. Poi stampa la somma.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    const MAX_RAND, TARGET = 11, 50
    var nRand int
    rand.Seed(time.Now().Unix())
    sommaRand := 0
    for sommaRand < TARGET {
        nRand = rand.Intn(MAX_RAND)
        fmt.Print(nRand, " ")
        sommaRand += nRand
    }
    fmt.Println()
    fmt.Println("Somma finale:", sommaRand)
}
