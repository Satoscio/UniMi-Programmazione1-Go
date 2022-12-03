/*
    Specifiche: Scrivere un programma frazioni.go che verifica se due frazioni
    num1/den1 e num2/den2 sono equivalenti e stampa "equivalenti" o
    "non equivalenti", a seconda del caso.
    Usare il tipo int per i dati in input.
    Trovate una soluzione che non faccia uso di float64.
*/

package main

import "fmt"

func main() {
    var num1, den1, num2, den2 int
    fmt.Print("Inserisci una frazione: ")
    fmt.Scan(&num1, &den1)
    fmt.Print("Inserisci una seconda frazione: ")
    fmt.Scan(&num2, &den2)

    if num1 * den2 == num2 * den1 {
        fmt.Println("Equivalenti")
    } else {
        fmt.Println("Non equivalenti")
    }
}
