/*
    Scrivere un programma es09.go che legge un numero n e stampa i numeri tra
    1 e n che sono dei quadrati.
*/

package main

import "fmt"

func main() {
    var n int
    fmt.Print("Inserisci un intero: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        if i*i <= n {
            fmt.Print(i*i, " ")
        } else {
            break
        }
    }
    fmt.Println()
}
