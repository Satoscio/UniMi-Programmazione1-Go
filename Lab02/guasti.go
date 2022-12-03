/*
    Specifiche: Scrivere un programma guasti.go che riceve in ingresso tre
    interi che segnalano eventuali guasti di un impianto. Per ciascun codice,
    se è diverso da 0, il programma deve stampare il guasto segnalato secondo
    la tabella che segue:

    codice      guasto
    cod. 1      caricamento acqua
    cod. 2      scarico acqua
    cod. 3      sistema di riscaldamento
*/

package main

import "fmt"

func main() {
    var g1, g2, g3 int

    fmt.Print("Inserisci i tre guasti: ")
    fmt.Scan(&g1, &g2, &g3)

    fmt.Print("Guasto 1: ")
    if g1 == 1 {
        fmt.Println("Caricamento acqua")
    } else if g1 == 2 {
        fmt.Println("Scarico acqua")
    } else if g1 == 3 {
        fmt.Println("Sistema di riscaldamento")
    }

    fmt.Print("Guasto 2: ")
    if g2 == 1 {
        fmt.Println("Caricamento acqua")
    } else if g2 == 2 {
        fmt.Println("Scarico acqua")
    } else if g2 == 3 {
        fmt.Println("Sistema di riscaldamento")
    }

    fmt.Print("Guasto 3: ")
    if g3 == 1 {
        fmt.Println("Caricamento acqua")
    } else if g3 == 2 {
        fmt.Println("Scarico acqua")
    } else if g3 == 3 {
        fmt.Println("Sistema di riscaldamento")
    }


}
