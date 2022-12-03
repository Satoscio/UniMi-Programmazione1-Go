/*
    Specifiche: Scrivere un programma Go mattino2.go che legga un orario
    (due numeri interi, ora e minuto) e stampi "mattino" se l'orario è compreso
    tra le 5:30 (incluso) e le 12:30 (escluso)
*/

package main

import "fmt"

func main() {
    var ora, minuti int
    fmt.Print("Inserire ora e minuti: ")
    fmt.Scan(&ora, &minuti)
    if ora < 5 || ora > 12 {
        fmt.Println("Non è mattino")
    } else if ora == 5 && minuti >= 30 {
        fmt.Println("Mattino")
    } else if ora == 12 && minuti < 30 {
        fmt.Println("Mattino")
    } else if ora > 5 && ora < 12 {
        fmt.Println("Mattino")
    } else {
        fmt.Println("Non è mattino")
    }
}
