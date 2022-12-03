/*
    Scrivere un programma crescente.go che legge da standard input una stringa
    di sole lettere minuscole e la stampa inserendo un '-' ogni volta che una
    lettera viene prima in ordine alfabetico della lettera precedente.

    Per esempio:
    data in input la parola ambire, il programma stampa am-bir-e
*/

package main

import "fmt"

func main() {
    var x string
    fmt.Print("Inserisci una parola: ")
    fmt.Scan(&x)

    fmt.Print(string(x[0]))

    for i := 1; i < len(x); i++ {

        if (x[i] < x[i-1]) {
            fmt.Print("-")
        }
        fmt.Print(string(x[i]))

    }
    fmt.Println()
}
