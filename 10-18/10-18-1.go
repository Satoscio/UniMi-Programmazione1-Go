package main

import "fmt"

func main() {
    var num, un, da, cn, totale int
    fmt.Print("Inserisci un numero con 3 cifre: ")
    fmt.Scan(&num)

    if num < 100 || num > 999 {
        fmt.Println("Numero non valido!")
    } else {
        un = num % 10
        da = (num / 10) % 10
        cn = (num / 100) % 10
        totale = un + da + cn

        if totale < 10 {
            fmt.Println("La somma delle tre cifre è minore di 10")
        } else {
            fmt.Println("La somma delle tre cifre è maggiore di 10")
        }
    }

}
