/*
    Leggere un testo (una sequenza di "parole", cioè di stringhe separate da
    spazi) fino a incontrare la parola "stop", e stampare il testo senza
    punteggiatura (ciao, come stai? -> ciao come stai)
*/

package main

import (
    "fmt"
    "unicode"
)

func main() {
    var s, s1 string
    parole := []string{}

    for {
        fmt.Scan(&s)
        if s == "stop" {
            break
        }
        parole = append(parole, s)
    }

    for _, k := range parole {
        for _, r := range k {
            if !unicode.IsPunct(r) {
                s1 += string(r)
            }
        }
        s1 += " "
    }

    fmt.Println(s1)
}
