/*
Problema: Scrivere un programma Go conversione_giorni.go che converta un numero
specificato di giorni (letto da tastiera) in anni, settimane, giorni senza mai usare l’operazione
di sottrazione. Si ignorino gli anni bisestili.
Nota. Usare costanti per il numero di giorni (365) in un anno e il numero di giorni in una
settimana (7).
Esempio di esecuzione:
numero di giorni da convertire?
1329
1329 giorni = 3 anni + 33 settimane + 3 giorni
*/

package main

import "fmt"

func main() {
    const G_ANNO = 365
    const G_SETT = 7
    var giorniInput, anni, sett, giorni int
    fmt.Print("Numero di giorni da convertire? ")
    fmt.Scan(&giorniInput)
    anni = giorniInput / G_ANNO
    sett = (giorniInput % G_ANNO) / G_SETT
    giorni = (giorniInput % G_ANNO) % G_SETT
    fmt.Println(giorniInput, "giorni =", anni, "anni +", sett, "settimane +", giorni, "giorni")
}
