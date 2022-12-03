/*
    Specifiche: Scrivere un programma Go mattino.go che legga un orario (intero)
    e stampi "mattino" se l'orario è compreso tra le 6 (incluso) e le 13 (escluso).
*/

package main

import "fmt"

func main() {
    var ora int
    fmt.Print("Inserire ora: ")
    fmt.Scan(&ora)
    fmt.Println(ora >= 6 && ora < 13)
}
