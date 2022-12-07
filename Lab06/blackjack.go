package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const N_SEMI, N_VALORI, N_CARTE = 4, 13, 52

var semi = [N_SEMI]string{"♥", "♦", "♣", "♠"}
var valori = [N_VALORI]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

type Carta struct {
	valore, seme string
}

func carta(n int) (Carta, bool) {
	var c1 Carta

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

func getValoreBJ(card Carta) int {
	switch card.valore {
	case "J", "Q", "K":
		return 10
	case "A":
		return 11
	default:
		x, _ := strconv.Atoi(card.valore)
		return x
	}
}

func mazzoPoker() (mazzo [N_CARTE]Carta) {
	for i := 0; i < N_CARTE; i++ {
		mazzo[i].valore = valori[i%N_VALORI]
		mazzo[i].seme = semi[i/N_VALORI]
	}
	return
}

func mischia(mazzo *[N_CARTE]Carta) {
	for i := 0; i < N_CARTE; i++ {
		n := rand.Intn(N_CARTE)
		mazzo[i], mazzo[n] = mazzo[n], mazzo[i]
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	x := estraiCarta()
	fmt.Println("Carta estratta:", x)
	fmt.Println("Valore della carta precedente:", getValoreBJ(x))
	fmt.Println("4 carte estratte:", dai4Carte())
	m := mazzoPoker()
	fmt.Print("\nMazzo:\n", m)
	mischia(&m)
	fmt.Print("\n\nMazzo mischiato:\n", m)
}
