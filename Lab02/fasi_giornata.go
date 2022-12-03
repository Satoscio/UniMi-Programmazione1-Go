/*
    Specifiche: Scrivere un programma fasi_giornata.go che che legga un orario
    (intero) e stampa la fase della giornata corrispondente, secondo la tabella
    che segue, o "orario non valido" se l'intero è fuori dall'intervallo 0-23

    ora         fase della giornata
    0 - 6       notte
    7 - 13      mattino
    14 - 18     pomeriggio
    19 - 23     sera
*/

package main

import "fmt"

func main() {
    var orario int
    fmt.Print("Inserisci un'ora 0-23: ")
    fmt.Scan(&orario)

    if orario < 0 || orario > 23 {
        fmt.Println("Orario non valido")
        return
    }

    if orario >= 0 && orario <= 6 {
        fmt.Println("Notte")
    } else if orario >= 7 && orario <= 13 {
        fmt.Println("Mattino")
    } else if orario >= 14 && orario <= 18 {
        fmt.Println("Pomeriggio")
    } else {
        fmt.Println("Sera")
    }
}
