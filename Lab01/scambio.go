/*
Problema: Scrivere un programma Go scambio.go che, dati in input due valori di tipo int,
controlli che i valori letti siano di tipo int, li stampi nell’ordine in cui sono stati forniti,
scambi i due valori (senza appoggiarsi a altre variabili), e poi esegua la stessa istruzione di
stampa (per mostrare che i valori sono stati effettivamente scambiati tra le due var).
Nota. Vedere nella documentazione di fmt.Scan come avere un rimando sull’esito della
lettura.
Esempio di esecuzione:
val1 e val2 (int): 10 20
errore: <nil>
prima dello scambio: 10 20
dopo lo scambio: 20 10
*/

package main

import "fmt"

func main() {
    var num1, num2 int
    fmt.Print("Val1 e Val2: ")
    _, errore := fmt.Scan(&num1, &num2)
    fmt.Println("Errore:", errore)
    fmt.Println("Prima dello scambio:", num1, num2)
    num1, num2 = num2, num1
    fmt.Println("Dopo lo scambio:", num1, num2)

}
