/*
	Scrivere un programma minu_maiu.go che legge da standard input
	una stringa e stabilisce se la stringa contiene solo minuscole,
	solo maiuscole o sia minuscole che maiuscole, quindi stampa un
	messaggio opportuno (es. "solo minuscole","solo minuscole",
	"sia minuscole che maiuscole").
*/

package main

import "fmt"

func main() {
	var s string
	var min, mai int
	fmt.Scan(&s)

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			min++
		} else {
			mai++
		}
	}

	if min == 0 && mai != 0 {
		fmt.Println("solo maiuscole")
	} else if min != 0 && mai == 0 {
		fmt.Println("solo minuscole")
	} else {
		fmt.Println("sia maiuscole che minuscole")
	}
}