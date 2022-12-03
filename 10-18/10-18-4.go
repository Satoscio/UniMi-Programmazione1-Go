// Dato n calcola 1+2+...+n senza usare la formula di Gauss

package main

import "fmt"

func main() {
    var n, somma int
    fmt.Print("> ")
    fmt.Scan(&n)

    for n > 0 {
        somma += n
        n--
    }

    fmt.Println(somma)
}
