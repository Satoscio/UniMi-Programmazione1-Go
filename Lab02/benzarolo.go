/*
    Specifiche: Scrivere un programma benzarolo.go che calcola il prezzo di un
    pieno, dati litri (float64) e tipo di carburante (int, 0=benzina, 1=diesel,
    2=alcol, 3=cherosene) secondo la tabella dei prezzi.
*/

package main

import "fmt"

func main() {
    const LITRO_BENZINA = 1.78
    const LITRO_DIESEL = 1.98
    const LITRO_ALCOOL = 1.2
    const LITRO_CHEROSENE = 1.1

    var carburante int
    var litri, costo_pieno float64
    fmt.Print("Inserisci il carburante e i litri: ")
    fmt.Scan(&carburante, &litri)

    if carburante == 0 {
        costo_pieno = LITRO_BENZINA * litri
    } else if carburante == 1 {
        costo_pieno = LITRO_DIESEL * litri
    } else if carburante == 2 {
        costo_pieno = LITRO_ALCOOL * litri
    } else if carburante == 3 {
        costo_pieno = LITRO_DCHEROSENE * litri
    } else {
        fmt.Println("Carburante non valido")
    }

    fmt.Println("Costo: €", costo_pieno)


}
