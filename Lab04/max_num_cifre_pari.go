/*
	Scrivere un programma max_num_cifre_pari.go che, data una sequenza di numeri
	(da leggere come stringhe), terminata dalla stringa "000", stampa il numero
	di cifre pari contenute nel numero che ne contiene più di tutti.
*/

package main

import "fmt"

func main() {
	var n, num, min, max int
	for {
		s, c := 0, 0
		fmt.Scan(&n)
		if n == 000 {
			break
		}
		temp := n
		for temp > 0 {
			s += temp % 10
			if s%2 == 0 {
				c++
			}
			temp /= 10
		}
		if c > min {
			max, min, num = c, c, n
		}
	}
	fmt.Printf("Numero con più cifre pari: %d con %d\n", num, max)
}
