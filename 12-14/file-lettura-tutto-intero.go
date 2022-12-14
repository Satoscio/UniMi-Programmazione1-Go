/*
	programma che legge da riga di comando il nome di un file
	e mette tutto il contenuto in una stringa sola
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	source := os.Args[1]
	b, _ := os.ReadFile(source)
	s := string(b)
	fmt.Println(s)
}