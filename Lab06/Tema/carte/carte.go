package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N_SEMI, N_VALORI, N_CARTE = 4, 13, 52

type Carta struct {
	valore, seme string
}

func carta(n int) (Carta, bool) {
	var c1 Carta
	var semi = [N_SEMI]string{"\u2665", "\u2666", "\u2663", "\u2660"}
	var valori = [N_VALORI]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	c1.seme = semi[n/N_VALORI]
	c1.valore = valori[n%N_VALORI]
	return c1, true
}

func estraiCarta() Carta {
	n := rand.Intn(N_CARTE)
	c1, _ := carta(n)
	return c1
}

func dai4Carte() [4]Carta {
	var carte [4]Carta
	for i := 0; i < 4; i++ {
		carte[i] = estraiCarta()
	}
	return carte
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(estraiCarta())
	fmt.Println(dai4Carte())
}
