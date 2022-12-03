//  date le due date di nascita di due persone, stabilire quale sia più vecchia

package main

import "fmt"

func main() {
    var anno1, mese1, giorno1, anno2, mese2, giorno2 int
    fmt.Print("Data di nascita di una prima persona in formato yyyy MM dd: ")
    fmt.Scan(&anno1, &mese1, &giorno1)
    fmt.Print("Data di nascita di una seconda persona in formato yyyy MM dd: ")
    fmt.Scan(&anno2, &mese2, &giorno2)

    if anno1 > anno2 {
        fmt.Println("La seconda persona è più grande")
    } else if anno1 == anno2 {
        if mese1 > mese2 {
            fmt.Println("La seconda persona è più grande")
        } else if mese1 == mese2 {
            if giorno1 > giorno2 {
                fmt.Println("La seconda persona è più grande")
            } else if giorno1 == giorno2 {
                fmt.Println("Sono nate nello stesso giorno!")
            } else {
                fmt.Println("La prima persona è più grande")
            }
        } else {
            fmt.Println("La prima persona è più grande")
        }
    } else {
        fmt.Println("La prima persona è più grande")
    }
}
