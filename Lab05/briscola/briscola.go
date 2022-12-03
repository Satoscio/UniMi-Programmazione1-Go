/*
Punti a briscola:
Asso (A): 11
3: 10
Re (K): 4
Donna (Q): 3
Fante (J): 2
7, 6, 5, 4, 2: 0
*/

package main

import "fmt"

func punti(s string) int {
    var somma int
    for i := 0; i < len(s); i++ {
        switch s[i] {
            case 'A':
                somma += 11
            case 'K':
                somma += 4
            case 'Q':
                somma += 3
            case 'J':
                somma += 2
            case '3':
                somma += 10
            default:
                somma += 0
        }
    }
    return somma
}

func main() {
    var mano string
    fmt.Print("Inserisci la tua mano di briscola:")
    fmt.Scan(&mano)

    fmt.Println("La tua mano ha", punti(mano), "punti")
}
