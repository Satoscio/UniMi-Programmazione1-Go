/*
    Scrivere un programma disegna_v.go che legge un intero positivo n e stampa
    una v di altezza n di asterischi.

    Esempio di esecuzione:
    dimensione v: 3

    *     *
     *   *
      * *
       *
*/

package main

import "fmt"

func main() {
    var x int
    fmt.Print("Dimensione V: ")
    fmt.Scan(&x)

    for i := 0; i < x; i++ {

        for h := 0; h < i; h++ {
            fmt.Print(" ")
        }

        fmt.Print("*")

        for h := 0; h < x-2; h++ {
            fmt.Print(" ")
        }

        fmt.Println("*")

    }

}
