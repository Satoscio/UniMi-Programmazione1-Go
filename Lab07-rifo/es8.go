/*
    Leggere un testo (una sequenza di "parole", cioè di stringhe separate da
    spazi) fino a incontrare EOF (ctrl D su una riga nuova) e stampare solo le
    parole di lunghezza maggiore della media delle lunghezze di tutte le parole
    del testo.
*/

package main

import "fmt"

func main() {
    var s string
    parole := []string{}
    var max, media int

    for {
        num, _ := fmt.Scan(&s) // num è 0 solamente quando c'è EOF (ctrl-D)
        if num == 0 {
            break
        }
        parole = append(parole, s)
    }

    for _, k := range parole {
        max += len(k)
    }
    fmt.Printf("\nMax: %d\n", max)
    media = max / len(parole)
    fmt.Printf("Media: %d\n", media)

    for _, k := range parole {

        if len(k) > media {
            fmt.Println(k)
        }

    }

}
