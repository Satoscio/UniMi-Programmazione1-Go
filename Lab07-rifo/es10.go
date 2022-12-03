/*
    Leggere un testo (una sequenza di "parole", cioè di stringhe separate da
    spazi) fino a incontrare EOF (ctrl D su una riga nuova) e stampare solo
    (e tutte) le parole che hanno lunghezza massima.
*/

package main

import "fmt"

func main() {
    var s string
    parole := []string{}
    var max int

    for {
        num, _ := fmt.Scan(&s)
        if num == 0 {
            break
        }

        if len(s) > max {
            max = len(s)
        }
        parole = append(parole, s)
    }
    fmt.Println("Parole con lunghezza max: ")
    for _, k := range parole {
        if len(k) == max {
            fmt.Println(k)
        }
    }

}
