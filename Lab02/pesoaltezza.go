/*
    Specifiche: Scrivere un programma pesoaltezza.go che riceve in ingresso due
    numeri (float64) rappresentanti peso in kg e altezza in cm di una persona e
    stampa un responso in funzione della tabella:

    peso        altezza     responso

    10 - 50     100 - 150   normopeso
    10 - 50     151 - 200   sottopeso
    51 - 100    100 - 150   sovrappeso
    51 - 100    151 - 200   normopeso
    101 +       100 - 150   molto sovrappeso
    101 +       151 - 200   sovrappeso

    (attenzione alle sovrapposizioni degli intervalli, decidere come risolverle)
*/

package main

import "fmt"

func main() {
    var peso, alt float64
    fmt.Print("Peso (kg) e altezza (cm): ")
    fmt.Scan(&peso, &alt)

    if peso >= 10 && peso <= 50 {
        if alt >= 100 && alt <= 150 {
            fmt.Println("Normopeso")
        } else {
            fmt.Println("Sottopeso")
        }
    } else if peso > 50 && peso <= 100 {
        if alt >= 100 && alt <= 150 {
            fmt.Println("Sovrappeso")
        } else {
            fmt.Println("Normopeso")
        }
    } else {
        if alt >= 100 && alt <= 150 {
            fmt.Println("Molto sovrappeso")
        } else {
            fmt.Println("Sovrappeso")
        }
    }
    
}
