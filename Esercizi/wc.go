// esercizio Word Counter della prova di esame del 24 gennaio 2023

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nb, nr, nw, nl int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		nl++
		for range scanner.Text() {
			nr++
		}
		nb += len([]byte(scanner.Text()))
	}

	fmt.Println("byte:", nb)
	fmt.Println("rune:", nr)
	fmt.Println("parole:", nw)
	fmt.Println("righe:", nl)
}
