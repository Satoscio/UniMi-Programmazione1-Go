/*
    Scrivi una funzione operazioni(n1, n2 int) (int, int, int) che accetta due
    interi e restituisce somma, prodotto e differenza.
    Scrivine una versione con variabili di ritorno nominate e una senza (puoi
    commentare una delle due versioni per caricare un file solo, altrimenti
    chiama i due file operazioni_a.go e operazioni_b.go)

    Scrivi un main per invocare e testare la funzione. Il programma legge da
    standard input due int. Chiama il programma operazioni.go e caricalo su upload
*/

package main

import "fmt"

/*
func operazioni(n1, n2 int) (int, int, int) {
    s := n1 + n2
    p := n1 * n2
    d := n1 / n2
    return s, p, d
}
*/

func operazioni(n1, n2 int) (s int, p int, d int) {
    s = n1 + n2
    p = n1 * n2
    d = n1 - n2
    return
}


func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)
    fmt.Println(operazioni(n1, n2))
}
