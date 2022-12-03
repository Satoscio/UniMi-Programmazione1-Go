// programma che trova divisori primi di un numero

package main

import "fmt"

func isPrime(n int) bool {
    for d := 2; d < n; d++ {
        if n % d == 0 {
            return false
        }
    }
    return true
}

/*  dato un numero come parametro, stampa i suoi divisori primi
    senza ripetizioni e separati da TAB, mettendo un a capo alla fine,
    e ne restituisce il numero
*/
func printAndCount(n int) (c int) {
    for d := 2; d < n; d++ {
        if n % d == 0 && isPrime(d) {
            fmt.Print(d, "\t") // \t è un TAB
            c++
        }
    }
    fmt.Println()
    return
}

func main() {
    var n int
    fmt.Print("Inserire un intero: ")
    fmt.Scan(&n)
    c := printAndCount(n)
    fmt.Println("Ci sono", c, "divisori")
}
