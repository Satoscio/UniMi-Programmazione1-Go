package main

import "fmt"

func main() {
    /*
        Scrivere un programma in Go che chieda 5 addendi interi e stampi la
        loro somma alla fine
    */
    const K = 5
	var n int
	s := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		s = s + n
	}
	fmt.Println(s)
}
