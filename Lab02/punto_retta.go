/*
    Specifiche: Scridvere un programma punto_retta.go che legge da standard input
    le cooordinate (x1,y1) di un punto P e i coefficienti m e q di una retta
    y = mx + q; controlla se il punto P sta sopra, sotto o appartiene alla
    retta y data, e stampa rispettivamente "sopra", "sotto", "sulla retta".
    Si dichiarino i dati in ingresso come float64.
    Nota Bene, trattando con numeri float, cosideriamo appartenenti alla retta
    i punti che distano dalla retta meno di 10^-6. In Go 10^-6 può essere
    scritto come 1e-06.
*/

package main

import "fmt"

func main() {
    const TOLLERANZA = 1e-06
    var x1, y1, m, q, y2, diff float64
    fmt.Print("Inserisci coordinate di un punto: ")
    fmt.Scan(&x1, &y1)
    fmt.Print("Inserisci i coefficienti 'm' e 'q' di una retta: ")
    fmt.Scan(&m, &q)

    y2 = (m * x1) + q
    
    if y1 > y2 {
        diff = y1 - y2
        if diff <= TOLLERANZA {
            fmt.Println("Il punto si trova sulla retta.")
        } else {
            fmt.Println("Il punto si trova sopra alla retta.")
        }
    } else if y2 > y1{
        diff = y2 - y1
        if diff <= TOLLERANZA {
            fmt.Println("Il punto si trova sulla retta.")
        } else{
            fmt.Println("Il punto si trova sotto alla retta.")
        }
    } else {
        fmt.Println("Il punto si trova sulla retta.")
    }
}
