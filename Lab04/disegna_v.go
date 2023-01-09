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
    x++
    s := x + (x-3)
    for i := 0; i < x; i++ {

        for i1 := 0; i1 < i; i1++ {
            fmt.Print(" ")
        }

        fmt.Print("*")

        for i2 := 0; i2 < s; i2++ {
            fmt.Print(" ")
        }
        s -= 2

        if i != x-1 {
            fmt.Println("*")
        } else {
            fmt.Println()
        }

    }

}
