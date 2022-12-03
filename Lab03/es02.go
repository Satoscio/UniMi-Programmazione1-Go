// Scrivere un programma es02.go che legge K = 5 numeri e di ciascuno stampa il doppio.

package main

import "fmt"

func main() {
    const K = 5
    for i := 0; i < K; i++ {
        n := 0
        fmt.Print("Inserisci un intero: ")
        fmt.Scan(&n)
        fmt.Println("Doppio:", n * 2)
    }
}
