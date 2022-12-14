/*
	Scrivere un programma posizioni_parole.go che legge una sequenza di stringhe
	da standard input e produce su standard output, per ogni stringa, la lista
	delle posizioni in cui essa compare nella sequenza (partendo dalla posizione 0).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	testo := []string{}
	lPos := map[string][]int{}
	for scan.Scan() {
		row := strings.Split(scan.Text(), " ")
		testo = append(testo, row...)
	}
	for i, k := range testo {
		lPos[k] = append(lPos[k], i)
	}
	fmt.Println(lPos)
}
