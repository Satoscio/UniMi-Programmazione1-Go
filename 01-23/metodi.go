package main

import "fmt"

type Rettangolo struct {
	base, altezza float64
}

func (r Rettangolo) stampa() {
	fmt.Println(r.base, r.altezza)
}

func main() {
	var r Rettangolo
	fmt.Print("base altezza: ")
	fmt.Scan(&r.base, &r.altezza)
	r.stampa()
}
