/*
    Specifiche: Scrivere un programma tasse.go che chiede reddito (int) e se
    coniugato (bool) e stampa le tasse da pagare secondo la seguente tabella.

    Usate costanti per le aliquote (10% e 25%) e per gli scaglioni (32000 e 64000).

    non coniugato       coniugato       %

    0 - 32000           0 - 64000       10%
    > 32000             > 64000         25%

*/

package main

import "fmt"

func main() {
    const ALIQ_1, ALIQ_2, SCAG_1, SCAG_2 = 10, 25, 32000, 64000

    var red int
    var con bool
    fmt.Print("Reddito? ")
    fmt.Scan(&red)
    fmt.Print("Coniugato? ")
    fmt.Scan(&con)

    if con == true {
        if red <= SCAG_2 {
            fmt.Println("Tasse:", ALIQ_1, "%")
        } else {
            fmt.Println("Tasse:", ALIQ_2, "%")
        }
    } else {
        if red <= SCAG_1 {
            fmt.Println("Tasse:", ALIQ_1, "%")
        } else {
            fmt.Println("Tasse:", ALIQ_2, "%")
        }
    }
}
