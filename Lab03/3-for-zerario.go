package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    /*
        Scrivere un programma in Go che generi interi [0, 20] casuali e li stampi.
        Il programma termina quando viene generato uno 0, stampa poi quanti
        interi ha generato prima dello 0
    */
    rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    numeriGenerati := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        numeriGenerati++
    }
    fmt.Println(numeriGenerati)
}
