/*
   Codice del programma estraiData
   Autore: Bogdan Budei
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var dataGMA string
	fmt.Print("data nel formato giorno/mese/anno: ")
	fmt.Scan(&dataGMA)
	g, m, a := stringGMA2GMA(dataGMA)
	fmt.Println("giorno", g)
	fmt.Println("mese", num2mese(m))
	fmt.Println("anno", a)

}

func stringGMA2GMA(data string) (int, int, int) {
	var g, m, a int
	data_sep := strings.Split(data, "/")
	fmt.Println(data_sep)
	g, _ = strconv.Atoi(data_sep[0])
	m, _ = strconv.Atoi(data_sep[1])
	a, _ = strconv.Atoi(data_sep[2])
	return g, m, a
}

func num2mese(m int) string {
	var mesi = [13]string{"", "gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"}
	return mesi[m]
}
