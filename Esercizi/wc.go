// esercizio Word Counter della prova di esame del 24 gennaio 2023
// Budei Bogdan

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nb, nr, nw, nl int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		nl++
		if line == "" {
			continue
		}
		for range line {
			nr++
		}
		nb += len([]byte(line))

		//  questa parte di codice mi è servita a stampare tutti i codepoint
		/*
			fmt.Printf("[%2.f] ", float64(nl))
			for _, r := range line {
				fmt.Printf("%d ", r)
			}
			fmt.Println()
		*/

		row := strings.Split(line, " ")
		for _, s := range row {
			if s == "" {
				continue
			}
			nw++
		}

	}
	nb += (nl - 1)
	fmt.Println("byte:", nb)
	fmt.Println("rune:", nr)
	fmt.Println("parole:", nw)
	fmt.Println("righe:", nl)
}
