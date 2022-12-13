/*
	Scrivere un programma contaLettere.go che legge un testo (da stdin)
	e stampa quante volte (anche 0) appare nel testo ciascuna lettera
	minuscola dell'alfabeto ('a'-'z').

	Il programma è dotato di una costante (const LEN_ALFABETO = 26)
	e di una funzione
	func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int)
	che, data una stringa s, aggiorna i conteggi delle lettere minuscole di s.
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

const LEN_ALFABETO = 26
const OFFSET = 97

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int) {
	for _, r := range s {
		if unicode.IsLower(r) {
			contaMinu[r-OFFSET]++
		}
	}
}

func main() {
	contaMinu := [LEN_ALFABETO]int{}
	lettereAl := [LEN_ALFABETO]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		aggiornaConteggi(scanner.Text(), &contaMinu)
	}
	for i := 0; i < LEN_ALFABETO; i++ {
		lettereAl[i] = string(i+97)
	}
	fmt.Println(lettereAl)
	fmt.Println(contaMinu)
}
