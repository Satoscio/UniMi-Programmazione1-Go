/*
    Scrivere un programma quadrati.go che legge da standard input un intero n
    positivo e calcola, utilizzando solo variabili di tipo int,
    il massimo quadrato (k^2) <= n.
*/

package main

import "fmt"

func main() {
    var n int
    fmt.Print("Inserisci un intero: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {

        if i*i <= n {

        } else {
            fmt.Println("Quadrato inferiore più grande:", (i-1)*(i-1))
            break
        }

    }

}
