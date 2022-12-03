/*
    Scrivere un programma es11.go che stampa ripetutamente "ore e minuti:"
    per chiedere due valori h e min fino a ottenere due valori validi,
    cioè h in [0,23] e min in [0,59], poi stampa "valori validi"
*/

package main

import "fmt"

func main() {
    var ore, min int

    for {
        fmt.Print("Inserisci ore e minuti: ")
        fmt.Scan(&ore, &min)
        if (ore >= 0 && ore <= 23) && (min >= 0 && min <= 59) {
            fmt.Println("Valori validi")
            break
        }
    }
}
