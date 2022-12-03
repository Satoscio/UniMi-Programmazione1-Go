package main

import "fmt"

func main() {

	const MENU = "Menu del giorno:\n" +
		"a. pizza\n" +
		"b. penne al pomodoro\n" +
		"c. cotoletta e patatine\n" +
		"d. crostata e caffe`" +
		"f. fine"

	fmt.Println(MENU)

	var scelta string

	for {
		fmt.Println("ordinazione?")
		fmt.Scan(&scelta)
		if "a" <= scelta && scelta <= "d" {
			break
		}
		fmt.Println("ordinazione non valida")
	}

	var ordinazione string
	switch scelta {
	case "a":
		ordinazione = "pizza"
	case "b":
		ordinazione = "penne al pomodoro"
	case "c":
		ordinazione = "cotoletta e patatine"
	case "d":
		ordinazione = "crostata e caffe`"
	}
	fmt.Println("hai ordinato", ordinazione)
}
