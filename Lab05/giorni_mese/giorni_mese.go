package main

import (
    "fmt"
    "strconv"
)

func giorniInMese(mese int) int {
    switch mese {
        case 2:
            return 28
        case 1, 3, 5, 7, 8, 10, 12:
            return 31
        default:
            return 30
    }
}

func main() {
    var data, mese string
    fmt.Print("Inserisci una data dd-MM-yyyy: ")
    fmt.Scan(&data)

    for i := 3; i < 5; i++ {
        mese += string(data[i])
    }

    s, _ := strconv.Atoi(mese)
    d := giorniInMese(s)

    fmt.Println("Il mese", s, "ha", d, "giorni")
}
