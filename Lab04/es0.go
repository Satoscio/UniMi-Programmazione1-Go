/*
    Scrivere un programma es0.go che fa le seguenti cose:

    legge un byte
    lo stampa
    stampa il byte precedente, il byte stesso, e quello successivo in ordine
    lessicografico (ASCII). Ad esempio, con 'd' come input, deve stampare: c d e

    stabilisce se è una lettera tra A e L, o altro (stampa "A-L" o "altro")
    poi legge una stringa e la stampa in verticale, una runa per riga.
*/

package main

import "fmt"

func main() {
    var x, y byte
    fmt.Print("Inserire un carattere: ")
    fmt.Scanf("%c", &x)
    fmt.Print("Inserire un altro carattere: ")
    fmt.Scanf("%c", &y)
    fmt.Println(string(x-1), string(x), string(x+1))


    fmt.Println(string(y))
    /*
    if y >= 'A' && y <= 'L' {
        fmt.Println("A-L")
    } else {
        fmt.Println("Altro")
    }
    */
}
