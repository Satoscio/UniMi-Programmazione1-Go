/*
   Scrivere un programma vocali.go che analizza un testo e conta le occorrenze
   delle vocali (sia minuscole che maiuscole, ma non le accentate) nel testo e
   stampa, ma solo per le vocali presenti nel testo, il numero di volte che le
   vocali stesse sono presenti nel testo.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contaVocali(s string, vocali map[rune]int) {
	var conta int
	voc := [10]rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	for i := 0; i < len(voc); i++ {
		conta = strings.Count(s, string(voc[i]))
		if conta != 0 {
			vocali[voc[i]] += conta
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conta := map[rune]int{}
	for scanner.Scan() {
		contaVocali(scanner.Text(), conta)
	}

	for i, s := range conta {
		fmt.Printf("Ci sono %d %c\n", s, i)
	}
}
