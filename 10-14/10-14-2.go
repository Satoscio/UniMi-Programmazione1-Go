// leggete due frazioni e dire quale è più grande

package main

import "fmt"

func main() {
    var nom1, nom2, den1, den2 int
    fmt.Print("Inserisci una frazione: ")
    fmt.Scan(&nom1, &den1)
    fmt.Print("Inserisci una seconda frazione: ")
    fmt.Scan(&nom2, &den2)
    
    if float64(nom1) / float64(den1) > float64(nom2) / float64(den2) {
        fmt.Println("La prima frazione è più grande")
    } else {
        fmt.Println("La seconda frazione è più grande")
    }
}
