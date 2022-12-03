package main

import "fmt"

func main() {
	/*
	   Scrivere un programma in Go che determini se un intero inserito sia
       positivo, negativo oppure nullo
	*/

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}
