/*
    Leggere un testo (una sequenza di "parole", cioè di stringhe separate da
    spazi), fino a incontrare la parola "stop", e stampare le parole in ordine
    alfabetico ("stop" esclusa).
*/

package main

import (
    "fmt"
)

func main() {
    var s string
    parole := []string{}

    for {
        fmt.Scan(&s)
        if s == "stop" {
            break
        }
        parole = append(parole, s)
    }

    fmt.Println()

    conta := len(parole) - 1
    it := 0

    for _, k := range parole { // itera le parole della slice

        for i := 0; i < len(k); i++ { // itera le lettere di ogni parola



        }

    }

    for i, k := range parole {

        for i2 := 0; i2 < len(k); i2++ {
            if k[i2] < it {
                it = 
            }
        }

    }

}
