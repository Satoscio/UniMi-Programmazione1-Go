/*
Scrivere un programma anagrammi.go che legge due stringhe da linea di
comando e valuta se le due stringhe sono una l'anagramma dell'altra
(la seconda stringa è formata da una permutazione dei caratteri della prima)
*/
package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
)

func isAnagram(s1, s2 string) bool {
	r1, r2 := []string{}, []string{}
	for _, r := range s1 {
		r1 = append(r1, string(r))
	}
	for _, r := range s2 {
		r2 = append(r2, string(r))
	}
	sort.Strings(r1)
	sort.Strings(r2)
	return reflect.DeepEqual(r1, r2)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Input errato")
		os.Exit(1)
	}
	if isAnagram(os.Args[1], os.Args[2]) {
		fmt.Println("Sono anagrammi")
	} else {
		fmt.Println("Non sono anagrammi")
	}
}
