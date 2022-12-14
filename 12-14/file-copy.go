/*
	programma che legge da riga di comando il nome di un file
	esistente e fa una copia del primo file sul secondo
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fileName1 := os.Args[1]
	fileName2 := os.Args[2]

	//source, err := strconv.Atoi(os.Args[2])
	
	source, err := os.Open(fileName1)
	defer source.Close()
	if err != nil {
		fmt.Println("Errore di lettura!")
		os.Exit(1)
	}

	dest, err := os.Create(fileName2)
	defer dest.Close()
	if err != nil {
		fmt.Println("Errore di apertura del file di destinazione!")
		os.Exit(1)
	}

	var buffer []byte
	buffer = make([]byte, 1024)

	for {
		nLetti, _ := source.Read(buffer)
		if nLetti == 0 {
			break
		}
		dest.Write(buffer[:nLetti])
	}
}