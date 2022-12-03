/*
Problema: Scrivere un programma Go no_out.go che dichiara una variabile int e una costante
= 10, assegna alla variabile un valore doppio della costante, ma non stampa niente (non
importa se non compila).
*/

package main

import "fmt"

func main() {
    var num1 int
    const num2 = 10
    num1 = num2 * 2
}
