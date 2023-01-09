/*
   Scrivere un programma fibonacci.go che legge un intero positivo n e stampa i
   numeri di fibonacci dal primo all'n-esimo, rappresentandoli come righe di
   asterischi, ciascuna lunga quanto il numero da rappresentare.
*/

package main

import "fmt"

func main() {
    var x int
    fmt.Print("Numero: ")
    fmt.Scan(&x)
    a, b := 1, 1

    fmt.Print("*\n*\n")

    for i := 2; i < x; i++{
        c := a + b
        a, b = b, c
        for i := 0; i < c; i++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
