package main

import "fmt"

func main() {
	/*
	   Scrivere un programma in Go che dato un numero intero in input, restituisca
       "Fizz" se il numero è divisibile per 3, "Buzz" se è divisibile per 5
       ed entrambi se soddisfa entrambi i criteri precedenti
	*/

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
