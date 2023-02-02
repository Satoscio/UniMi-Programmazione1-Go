/*
	Scrivere un programma palindromo.go, dotato di:

	una funzione ricorsiva isPalindrome(s string) bool che stabilisca se una stringa è palindroma
	una funzione main() che, data una stringa come argomento sulla linea di comando, emetta nel
	flusso d'uscita il messaggio "s è palindroma", se s
	è palindroma, e "s non è palindroma" altrimenti.

	La stringa vuota e le stringhe di lunghezza 1 sono considerate palindrome.
*/

package main

import (
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return isPalindrome(s[1 : len(s)-1])
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("arg sbagliati")
		os.Exit(1)
	}
	s := os.Args[1]
	if isPalindrome(s) {
		fmt.Printf("%s è palindroma", s)
	} else {
		fmt.Printf("%s non è palindroma", s)
	}
}
