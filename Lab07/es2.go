/*
    Problema 2.
    Leggere un testo (una sequenza di "parole", cioè di stringhe separate da
    spazi) e stampare una parola per riga, partendo dall'ultima.
*/

package main

import (
    "fmt"
    "os"
    "bufio"
)

func itera(s string) {

}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    s := ""

    for scanner.Scan() {
        s += scanner.Text()
    }
    m := itera(os.Args[1])
    printMap(m)
}
