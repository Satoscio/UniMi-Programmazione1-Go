/*
    Specifiche: Scrivere un programma tariffe.go che chiede all’utente l’età (int)
    e se è studente (bool) e stampa il costo del biglietto di ingresso secondo
    la tabella che segue.
    Provare a risolvere l'esercizio utilizzando solo 4 (e non di più) istruzioni
    di assegnamento a una variabile tariffa e una sola istruzione di stampa
    finale per la stampa del costo.
*/
package main

import "fmt"

func main() {
    var eta int
    var tariffa float64
    var studente bool
    fmt.Print("Quanti anni hai? ")
    fmt.Scan(&eta)
    fmt.Print("Sei uno studente? (t/f): ")
    fmt.Scan(&studente)

    if e4ta < 8 {
        tariffa = 0.0
    } else if (eta >= 9 && eta < 18) || studente == true {
        tariffa = 5.0
    } else if eta >= 18 || eta < 26 {
        tariffa = 7.5
    } else {
        tariffa = 10.0
    }

    fmt.Println("La tariffa è:", tariffa, "euro")

}
