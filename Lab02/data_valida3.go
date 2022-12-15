/*
	Specifiche: Scrivere una terza versione data_valida3.go
	che legge tre interi gg, mm e aaaa (anno) e verifica che
	sia una data valida. Questa versione finale tiene conto
	anche degli anni bisestili.
*/

package main

import "fmt"

func main() {

	var gg, mm, anno int
	var bis bool
	fmt.Scan(&gg, &mm, &anno)

	if anno % 4 == 0 {
		if anno % 100 == 0 {
			if anno % 400 == 0 {
				bis = true
			} else {
				bis = false
			}
		} else {
			bis = true
		}
	} else {
		bis = false
	}

	switch mm {
		case 1, 3, 5, 7, 8, 10, 12:
			if gg < 1 || gg > 31 {
				fmt.Println("Data non valida")
			} else {
				fmt.Println("Data valida")
			}
		case 4, 6, 9, 11:
			if gg < 1 || gg > 30 {
				fmt.Println("Data non valida")
			} else {
				fmt.Println("Data valida")
			}
		case 2:
			if bis {
				if gg < 1 || gg > 29 {
					fmt.Println("Data non valida")
				} else {
					fmt.Println("Data valida")
				}
			} else {
				if gg < 1 || gg > 28 {
					fmt.Println("Data non valida")
				} else {
					fmt.Println("Data valida")
				}
			}
			
		default:
			fmt.Println("Data non valida")
	}
	

}