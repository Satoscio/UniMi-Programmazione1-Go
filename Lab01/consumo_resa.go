/*
    Problema. Scrivere un programma Go consumo_resa.go che calcola il consumo
    medio e la resa di un motore date la distanza totale percorsa (in km) e
    la quantità di carburante utilizzata (in litri). I valori in input sono di
    tipo float64.

    Nota. Il consumo medio di carburante si esprime in l/km ed è la quantità di
    carburante che occorre in media per percorrere un km di strada. La resa di
    un motore è data dalla distanza percorsa in media con un litro di carburante
    e si esprime in km/l.
*/

package main

import "fmt"

func main() {
    var dist, carb float64
    fmt.Print("Distanza percorsa in km: ")
    fmt.Scan(&dist)
    fmt.Print("Litri di carburante utilizzati: ")
    fmt.Scan(&carb)

    fmt.Println("Consumo medio:", carb/dist, "l/km")
    fmt.Println("Resa del motore:", dist/carb, "km/l")
}
