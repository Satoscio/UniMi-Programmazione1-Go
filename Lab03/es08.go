/*
    Scrivere un programma es08.go che legge un numero n e stampa la tabellina
    di n (cioè n*1, n*2, ..., n*10)
*/

package main

import "fmt"

func main() {
    var n int
    const TAB_MAX = 10
    fmt.Print("Inserisci un intero: ")
    fmt.Scan(&n)

    for i := 1; i <= TAB_MAX; i++ {
        fmt.Println("n *", i, "=", n*i)
    }
}
