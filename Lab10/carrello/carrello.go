package main

import (
	"fmt"
	"strconv"
	"strings"
)

const PMAX = 15

type Carrello struct {
	pos, carico int
}

func (c Carrello) String() string {
	return fmt.Sprintf("Carrello in pos %d con carico %d\n", c.pos, c.carico)
}

func aggiornaStato(c *Carrello, posizione, carico int) bool {
	if posizione < 0 || carico < 0 {
		return false
	}
	c.pos = posizione
	c.carico += carico
	return true
}

func parseString(s string) []int {
	s = strings.ReplaceAll(s, " ", "0")
	sl := strings.Split(s, "|")
	out := []int{}
	for _, el := range sl {
		temp, _ := strconv.Atoi(el)
		out = append(out, temp)
	} 
	return out
}

func reverseParse(c Carrello, sl []int) {
	fmt.Print("|")
	for _, n := range sl {
		if n != 0 {
			fmt.Print(n, "|")
		} else {
			fmt.Print(" |")
		}
	}
	fmt.Println()
	fmt.Print(c.String())
}

func main() {
	var c Carrello
	var count, peso_max int
	listaPesi := map[int]int{}
	// gestione input WIP
	sl := parseString("| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |")
	// reverseParse(c, sl)
	for c.pos = 0; c.pos < len(sl); c.pos++ {

		if sl[c.pos] != 0 {

			if sl[c.pos] > peso_max {
				peso_max = sl[c.pos]
			}

			if c.carico + sl[c.pos] < PMAX {
				aggiornaStato(&c, c.pos, sl[c.pos])
				listaPesi[sl[c.pos]]++
				sl[c.pos] = 0
			} else {
				reverseParse(c, sl)
				aggiornaStato(&c, 0, 0)
				sl[0] += c.carico
				c.carico = 0
				count++
			}
		}

	}

	reverseParse(c, sl)
	aggiornaStato(&c, 0, 0)
	sl[0] += c.carico
	c.carico = 0
	reverseParse(c, sl)
	fmt.Printf("Numero viaggi: %d\n", count)
	fmt.Printf("Peso max: %d\n", peso_max)
	fmt.Println("Oggetti trovati:")
	for k, v := range listaPesi {
		fmt.Printf("%d di peso %d\n", k, v)
	}
}