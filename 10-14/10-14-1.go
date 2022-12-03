package main

import "fmt"

func main() {
    var annoCorr, annoNasc int
    fmt.Print("In che anno sei nato? ")
    fmt.Scan(&annoNasc)
    fmt.Print("In che anno siamo? ")
    fmt.Scan(&annoCorr)
    if annoCorr - annoNasc >= 18 {
        fmt.Println("Sei maggiorenne")
    } else {
        fmt.Println("Sei minorenne")
    }
}
