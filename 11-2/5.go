package main

import (
    "fmt"
    "unicode"
)

func main() {
    var x string
    fmt.Scan(&x)
    n := 0

    for _, r := range x { // devo mettere _ se non mi serve usare l'indice
        /*if r >= 'a' && r <= 'z' {
            n++
        } */
        if unicode.IsLower(r) {
            n++
        }
    }
    fmt.Println("Ci sono", n, "lettere minuscole")
}
