// calcolare giorni da 1/1/1970 a una data inserita

package main

import "fmt"

func main() {
    var d, m, a int
    fmt.Print("Inserisci la tua data di nascita: ")
    fmt.Scan(&d, &m, &a)
    birthdate := daysFromEpoch(d, m, a)
    fmt.Println("Sono passati", birthdate, "giorni dal 1/1/1970")

    fmt.Print("Inserisci la data di oggi: ")
    fmt.Scan(&d, &m, &a)
    now := daysFromEpoch(d, m, a)
    fmt.Println("Sono passati", now, "giorni dal 1/1/1970")

    age := now - birthdate

    fmt.Println("Hai", age, "giorni")
}

func isLeap(a int) bool {
    return a % 400 == 0 || (a % 4 == 0) && (a % 100 != 0)
}

func daysFromEpoch(d, m, a int) int {
    var res int

    for i := 1970; i < a; i++ {
        if isLeap(i) {
            res += 366
        } else {
            res += 365
        }
    }

    for i := 1; i < m; i++ {
        if i == 11 || i == 4 || i == 6 || i == 9 {
            res += 30
        } else if i == 2 {
            if isLeap(i) {
                res += 29
            } else {
                res += 28
            }
        } else {
            res += 31
        }
    }
    res += d
    return res
}
