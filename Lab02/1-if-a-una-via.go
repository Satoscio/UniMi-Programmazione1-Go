package main

import "fmt"

func main() {
	/*
	   Scrivere un programma in Go che verifichi che avvisi se il voto inserito
       sia un voto non valido (al di fuori dell'insieme [0,30])
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
