/*
	programma che legge da riga di comando il nome di un file
	e stampa una riga per volta
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	fileName := os.Args[1]
	
	source, err := os.Open(fileName)
	defer source.Close()

	if err != nil {
		fmt.Println("Errore di lettura!")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(source)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}