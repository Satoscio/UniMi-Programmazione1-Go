/*
    Problema. Scrivere un programma Go intersezione_rette.go che, date le
    equazioni di due rette, stabilisce in che punto si intersecano. I dati in
    input sono di tipo float64.

    Nota. L’equazione della retta è: y = mx + q. Dal punto di vista matematico,
    occorre risolvere un sistema di primo grado.

    Esempio di esecuzione:

    retta 1: m e q? 1 4
    retta 2: m e q? 2 6
    intersezione in (-2,2)
*/

package main

import "fmt"

func main() {
    var m1, m2, q1, q2 float64
    fmt.Print("Inserisci m e q della retta 1: ")
    fmt.Scan(&m1, &q1)
    fmt.Print("Inserisci m e q della retta 2: ")
    fmt.Scan(&m2, &q2)

    xi := (q2 - q1) / (m1 - m2)
    yi := m1 * xi + q1

    fmt.Println("Intersezione:", xi, yi)

}
