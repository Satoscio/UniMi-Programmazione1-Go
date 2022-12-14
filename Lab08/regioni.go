/*
	Scrivere un programma regioni.go che che legge da standard input (usate la
	ridirezione dal file capoluoghi.csv fornito, lo trovate qui) una serie di
	informazioni sui capoluoghi organizzate su righe, nel seguente formato:

	Nome,Sigla,Regione,Popolazione,Superficie,Densita,Altitudine

	e memorizza nome e regione in modo che sia possibile ottenere la stampa della
	lista dei capoluoghi delle regioni i cui nomi sono stati forniti da linea di comando.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//args := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	reg := map[string]string{}

	for scanner.Scan() {
		all := strings.Split(scanner.Text(), ",")
		reg[all[0]] = all[2]
	}

	fmt.Println(reg)
	/*
		for _, r := range args {
			fmt.Println(reg[r])
		}
	*/
}
