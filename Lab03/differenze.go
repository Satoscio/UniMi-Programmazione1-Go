/*
    Scrivere un programma differenze.go che legge una serie di valori float64
    da tastiera e stampa le differenze, cioè la differenza tra il secondo e il
    primo, tra il terzo e il secondo, e così via.
    Il programma termina quando riceve in input il valore 0.
*/

package main

import "fmt"

func main() {
    var x, y float64
    for {
        fmt.Print("Due float: ")
        fmt.Scan(&x, &y)
        if x == 0.0 || y == 0.0 {
            break
        }
        fmt.Println(y-x)
        
    }
}
