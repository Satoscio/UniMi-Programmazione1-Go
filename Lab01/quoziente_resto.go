/*
Problema: Scrivere un programma Go quoziente_resto.go che legge da standard input un
dividendo e un divisore (interi), calcola il quoziente e il resto e li stampa.
Nota. L’operatore per la divisione (/) tra interi calcola la parte intera del risultato; l’operatore
per il resto della divisione è %.
Esempio di esecuzione:
dividendo: 2500
divisore: 235
quoziente = 10
resto = 150
*/

package main

import "fmt"

func main() {
    var dividendo, divisore, quoziente, resto int
    fmt.Print("Dividendo: ")
    fmt.Scan(&dividendo)
    fmt.Print("Divisore: ")
    fmt.Scan(&divisore)
    quoziente = dividendo / divisore
    resto = dividendo % divisore
    fmt.Println("Quoziente =", quoziente)
    fmt.Println("Resto =", resto)
}
