/*
    Scrivere un programma euclide.go che legge da standard input due interi a e b,
    con a >= b, e calcola il MCD tra i due numeri con l'algoritmo di Euclide.

    Algoritmo di Euclide:
    Dati due numeri naturali a e b:
    1. controlla se b è zero.
    2. se lo è, a è il MCD.
    3. se non lo è, assegna ad r il resto della divisione a / b,
       poni a = b e b = r e ripeti da 1.
*/

package main

import "fmt"

func main() {
    var a, b, r int
    fmt.Print("Inserisci due interi (il primo deve essere maggiore): ")
    fmt.Scan(&a, &b)

    for {
        if b == 0 {
            fmt.Println("Il MCD è", a)
            break
        }
        r = a % b
        a, b = b, r
    }
}
