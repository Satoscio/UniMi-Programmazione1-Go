/*
	Scrivere un programma sommaRicorsiva.go, dotato di:

	una funzione ricorsiva func recursiveSum(list []int) int
	che calcola la somma degli interi della slice list passata come parametro,
	una funzione main() che legga da standard input (ctrl D per terminare) una lista di numeri
	interi ed emetta nel flusso d'uscita la somma dei numeri letti.

	La somma di una lista vuota è 0.
*/

package main

import "fmt"

func recursiveSum(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + recursiveSum(list[1:])
}
func main() {
	list := []int{2, 3, 56, 90}
	fmt.Println(recursiveSum(list))

}
