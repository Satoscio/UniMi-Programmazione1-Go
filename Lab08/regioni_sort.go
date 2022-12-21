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
	"sort"
	"strings"
)

func unique(s []string) []string {
    inResult := make(map[string]bool)
    var result []string
    for _, str := range s {
        if _, ok := inResult[str]; !ok {
            inResult[str] = true
            result = append(result, str)
        }
    }
    return result
}

func main() {
	args := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	reg := map[string]string{}

	scanner.Scan()
	for scanner.Scan() {
		all := strings.Split(scanner.Text(), ",")

		for range all {
			reg[all[0]] = all[2]
		}
	}
	/*
	for z, r := range reg {
		for i := range args {
			if r == args[i] {
				fmt.Println(reg[z], z)
			}
		}
	}
	*/
	key := []string{}
	val := []string{}

	for k, v := range reg {
		key = append(key, k)
		val = append(val, v)
	}
	val = unique(val)

	sort.Strings(args)
	sort.Strings(key)
	sort.Strings(val)

	for _, z := range args {
		for _, v := range key {
			if reg[v] == z {
				fmt.Printf("%s: %s\n", reg[v], v)
			}
		}
		fmt.Println()
	}
}
