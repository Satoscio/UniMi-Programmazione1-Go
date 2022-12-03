/*
    Scrivere un programma indovina_numero.go che chiede all'utente di indovinare
    un numero intero random x tra 1 e MAX, (dove MAX è una costante definita nel
    programma) e ripete la richiesta fino a che l'utente indovina, oppure fino a
    un massimo di MAX/2 tentativi.

    Il programma stampa il numero di tentativi che sono stati necessari per
    indovinare oppure il messaggio "hai perso, il numero era x".
    Se il numero digitato dall'utente è fuori dall'intervallo [1,MAX], il
    tentativo non viene considerato e viene visualizzato il messaggio "fuori
    intervallo!", senza interrompere l'esecuzione.

    Utilizzare la funzione rand.Intn del package "math/rand" per fissare il
    numero da indovinare.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    var n int
    const MAX = 10
    rand.Seed(time.Now().Unix())
    numCasuale := rand.Intn(MAX) + 1
    tent := 1
    for i := 0; i < MAX/2; i++ {

        fmt.Print("Indovina il numero [1,10]: ")
        fmt.Scan(&n)

        if n == numCasuale {
            fmt.Println("Indovinato in", tent, "tentativi")
            break
        } else if n > 10 || n < 1 {
            fmt.Println("Fuori intervallo!")
            i--
        } else {
            fmt.Println("Non hai indovinato")
            tent++
        }
        
    }


}
