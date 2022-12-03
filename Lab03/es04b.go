/*
    Scrivere un programma es04b.go che genera K = 10 numeri casuali in [0,10],
    conta quanti sono pari e stampa questo risultato.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    var nRand, sommaPari int
    const MAX_RAND, K = 11, 10
    rand.Seed(time.Now().Unix())
    for i := 0; i < K; i++ {
        nRand = rand.Intn(MAX_RAND)
        fmt.Print(nRand, " ")
        if nRand % 2 == 0 {
            sommaPari += nRand
        }
    }
    fmt.Println()
    fmt.Println("Somma dei numeri pari:", sommaPari)
}
