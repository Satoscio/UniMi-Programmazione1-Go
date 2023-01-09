/*
	Scrivere un programma vicino.go che, data una serie
	di 5 valori interi tra 0 e 20, trova il valore più
	vicino a TARGET, dove TARGET è una costante definita
	nel programma.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const TARGET = 7
	var a int
	min := int(^uint(0) >> 1)
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		diff := int(math.Abs(float64(a - TARGET)))
		fmt.Println(diff)
		if diff < min {
			min = diff
		}
	}
	fmt.Println("Più vicino:", min+TARGET)
}
