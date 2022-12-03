/*
Problema: Scrivere un programma Go echo.go che legge un valore di tipo int da tastiera e lo
stampa.
*/

package main

import "fmt"

func main() {
    var num int
    fmt.Scan(&num)
    fmt.Println(num)
}
