/*
    Problema. Scrivere un programma Go conversione_secondi.go che converta un
    numero dato di secondi (fornito dall’utente) in giorni, ore, minuti, secondi
    senza mai usare l’operazione di sottrazione.

    Esempio di esecuzione:

    numero di secondi: 123456
    g:h:m:s = 1:10:17:36

    secondi in un giorno: 86400
    secondi in un'ora   : 3600
    secondo in un minuto: 60
*/

package main

import "fmt"

func main() {
    const SEC_GIO, SEC_ORA, SEC_MIN = 86400, 3600, 60
    var secIn, gio, ore, min, secOut int
    fmt.Print("Inserisci dei secondi: ")
    fmt.Scan(&secIn)

    gio = secIn / SEC_GIO
    ore = (secIn % SEC_GIO) / SEC_ORA
    min = ((secIn % SEC_GIO) % SEC_ORA) / SEC_MIN
    secOut = ((secIn % SEC_GIO) % SEC_ORA) % SEC_MIN

    //fmt.Printf("gg:hh:mm:ss: %d:%d:%d:%d\n", gio, ore, min, secOut)
    fmt.Println("gg:hh:mm:ss:", gio, ore, min, secOut)
}
