/*
    Leggere un testo (una sequenza di "parole", cioè di stringhe separate da
    spazi) fino a incontrare la parola "stop", e stampare una parola per riga,
    partendo dall'ultima ("stop" esclusa).
*/

package main

import (
    "fmt"
)

func main() {
    var s string

    for {
        fmt.Scan(&s)
        if s == "stop" {
            break
        }
        fmt.Println(s)
    }

}
