package main

import (
	"fmt"
)

type Segmento struct {
	estremi, interno byte
	orizzontale bool
	lunghezza int
}

func (seg Segmento) String() (dis string) {
	dis += string(seg.estremi)

	for i := 0; i < seg.lunghezza-2; i++ {
		if !seg.orizzontale {
			dis += "\n"
		}
		dis += string(seg.interno)
	}

	if !seg.orizzontale {
		dis += "\n"
	}

	dis += string(seg.estremi)
	return
}

func (x *Segmento) allunga(n int) {
	x.lunghezza += n
}

func main () {
	var seg Segmento

	fmt.Print("Inserisci in ordine: estremi, interni, orizzontale(t/f) e lunghezza: ")
	fmt.Scanf("%c %c %t %d", &seg.estremi, &seg.interno, &seg.orizzontale, &seg.lunghezza)
	fmt.Println(seg.String())
	seg.allunga(5)
	fmt.Println("\n", seg.String())
}