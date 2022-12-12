/*
	Scrivere un programma giornoSettimana.go, dotato di un array giorniDellaSettimana
	[7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
	che legge da linea di comando il nome del giorno della settimana
	del 1° gennaio di un dato anno. Se il giorno non è fra
	"lun", "mart", "merc", "giov", "ven", "sab", "dom", il programma
	avvisa e termina. Altrimenti poi accetta dall'utente dei numeri
	interi (tra 1 e 365), corrispondenti a giorni di quell'anno, e
	per ciascuno stampa il giorno della settimana corrispondente.
	Il programma termina quando l'utente digita -1.

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var trovato bool
	var indGiorno, giorno, sub int
	giornoMag := 1
	giorniDellaSettimana := [7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
	for i, d := range giorniDellaSettimana {
		if os.Args[1] == d {
			indGiorno = i
			trovato = true
		}
	}
	if !trovato {
		fmt.Println("Non c'è!")
		os.Exit(1)
	}
	
	for {
		giornoMag = 1
		fmt.Print("Inserisci un giorno [1, 365]: ")
		fmt.Scan(&giorno)
		if giorno == -1 {
			fmt.Println("Fine programma.")
			break
		}

		for giornoMag <= giorno {
			giornoMag += 7
		}
		sub = indGiorno - (giornoMag - giorno)
		if sub < 0 {
			sub = 7 + sub
		}
		fmt.Printf("Il giorno %d è %s\n", giorno, giorniDellaSettimana[sub])
	}


}