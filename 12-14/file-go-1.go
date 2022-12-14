/*
	programma che legge da riga di comando il nome di un file e un numero e stampa
	quel numero di bytee del file
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := os.Args[1]
	nByte, _ := strconv.Atoi(os.Args[2])
	
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println("Errore di lettura!")
		os.Exit(1)
	}

	var b []byte
	b = make([]byte, nByte)

	x, err := f.Read(b)

	fmt.Printf("Ho letto %d byte\n", x)
	fmt.Println("Byte letti:", b)
	fmt.Println("Stringa corrispondente:", string(b))
}