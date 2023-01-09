/*
	Si vuole scrivere un programma max_somma_cifre.go che

	legge da standard input una serie di numeri >= 0, terminata da 999,
	trova il numero (escludendo 999) la cui somma delle cifre è la maggiore
	e stampa tale somma.
*/

package main

import "fmt"

func main() {
	var n, min, maxSum, num int
	for {
		s := 0
		fmt.Scan(&n)
		if n == 999 {
			break
		}
		temp := n
		for temp > 0 {
			s += temp % 10
			temp /= 10
		}
		if s > min {
			maxSum, min, num = s, s, n
		}
	}
	fmt.Printf("Somma max %d nel numero %d\n", maxSum, num)
}
