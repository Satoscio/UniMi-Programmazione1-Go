// dato un num intero stampare la somma delle sue cifrepackage main
package main

import "fmt"

func main() {
    var num, tot int
    fmt.Print("Inserisci un numero intero: ")
    fmt.Scan(&num)

    for num != 0 {
        tot += (num % 10)
        num /= 10
    }

    fmt.Println(tot)
}
