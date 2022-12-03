/*
    Scrivere un programma es05.go che legge un numero n e stampa i numeri da 1 a n.
*/

package main

import "fmt"

func main() {
    var n int
    fmt.Print("Inserisci un intero: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        fmt.Print(i, " ")
    }

    fmt.Println()
}
