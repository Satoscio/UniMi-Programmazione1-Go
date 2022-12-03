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

    y := 1

    for i := 0; i < x; i++{
        fmt.Println(y)
        y += y

    }

    y, z := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
}
