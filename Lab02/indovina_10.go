/*
    Specifiche: Scrivere un programma indovina_10.go che fissa nel programma
    un numero intero (a vostra scelta) tra 1 e 10 da indovinare, legge un
    intero da standard input e:

    - se il numero in input è fuori dall’intervallo [1-10],
    stampa “Valore non valido”;

    - se il numero è quello fissato, stampa “Hai indovinato!”;

    - altrimenti stampa “Non hai indovinato”.
*/

package main

import "fmt"

func main() {
    const NUMERO_SEGRETO = 7
    var num int
    fmt.Print("Inserisci un numero da 1 a 10: ")
    fmt.Scan(&num)

    if num > 10 || num < 1 {
        fmt.Println("Valore non valido")
    } else if num == NUMERO_SEGRETO {
        fmt.Println("Hai indovinato!")
    } else {
        fmt.Println("Non hai indovinato")
    }
}
