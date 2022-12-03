/*
    Scrivere un programma disegna_slash.go che legge un intero positivo n e
    stampa un backslash (\) di altezza n composto da asterischi.

    Esempio di esecuzione:
    dimensione \: 3

    *
     *
      *
*/

package main

import "fmt"

func main() {
    var x int
    fmt.Print("Dimensione \\: ")
    fmt.Scan(&x)

    for i := 0; i < x; i++ {

        for h := 0; h < i; h++ {
            fmt.Print(" ")
        }

        fmt.Println("*")
    }
    
}
