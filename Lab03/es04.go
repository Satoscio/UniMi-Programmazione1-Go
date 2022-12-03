/*
    Scrivere un programma es04.go che genera K = 10 numeri casuali in [0,10]
    e li stampa su una riga, separati da spazi.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    var nRand int
    const MAX_RAND, K = 11, 10
    rand.Seed(time.Now().Unix())
    for i := 0; i < K; i++ {
        nRand = rand.Intn(MAX_RAND)
        fmt.Print(nRand, " ")
    }
    fmt.Println()
}
