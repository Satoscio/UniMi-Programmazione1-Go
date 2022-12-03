/*
Problema: Scrivere un programma Go operazioni.go che, dati due numeri float64 da standard
input, esegua addizione, sottrazione, moltiplicazione e divisione dei due numeri e stampi i
risultati ottenuti.
Nota: le operazioni possono essere fatte direttamente nelle istruzioni di stampa.
*/

package main

import "fmt"

func main() {
    var num1, num2 float64
    fmt.Print("Inserisci due float: ")
    fmt.Scan(&num1, &num2)
    fmt.Println("Somma:", num1 + num2)
    fmt.Println("Sottrazione:", num1 - num2)
    fmt.Println("Moltiplicazione:", num1 * num2)
    fmt.Println("Divisione:", num1 / num2)

}
