package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    /*
        Scrivere un programma in Go che generi degli interi casuali e li sommi
        fino a quando la somma non è maggiore di 20
    */
    const TARGET = 20
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n int

    t := 0
    for somma < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        somma += n
    }
    fmt.Println(somma)
}
