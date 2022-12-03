/*
    Specifiche: Scrivere un programma intorno.go che legge due float64 x e y
    e stampa "vicino all'origine" se il punto di coordinate (x,y) è abbastanza
    vicino (a meno di 10^-5) all'origine (0,0),  mentre stampa "non vicino
    all'origine" in caso contrario.
*/

package main

import "fmt"

func main() {
    const EPSILON = 1e-5
    var x, y float64
    fmt.Print("Inserisci coordinate di un punto: ")
    fmt.Scan(&x, &y)

    if x <= EPSILON && y <= EPSILON {
        fmt.Println("Vicino all'origine")
    } else {
        fmt.Println("Non vicino all'origine")
    }
}
