package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Capoluogo struct {
	nome, sigla, regione                         string
	popolazione, superficie, densita, altitudine int
}

func main() {
	capl := map[string]Capoluogo{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // serve a saltare la riga dell'intestazione

	for scanner.Scan() {
		var row Capoluogo
		all := strings.Split(scanner.Text(), ",") // slice di string splitatte della riga del file .csv
		als := all[:3]                            // prendo la parte in string
		aln := all[3:]                            // prendo la parte intera
		nms := []int{}                            // creo una slice di interi vuota

		for _, i := range aln { // popolo 'nms' convertendo i valori di 'aln' in int
			j, _ := strconv.Atoi(i)
			nms = append(nms, j)
		}
		// associo a ogni campo della struct il proprio valore
		row.nome = als[0]
		row.sigla = als[1]
		row.regione = als[2]
		row.popolazione = nms[0]
		row.superficie = nms[1]
		row.densita = nms[2]
		row.altitudine = nms[3]
		capl[row.sigla] = row
	}

	sig := os.Args[1]
	fmt.Printf("Nome :\t %s\n", capl[sig].nome)
	fmt.Printf("Sigla:\t %s\n", capl[sig].sigla)
	fmt.Printf("Reg. :\t %s\n", capl[sig].regione)
	fmt.Printf("Abit.:\t %d\n", capl[sig].popolazione)
	fmt.Printf("Area :\t %d\n", capl[sig].superficie)
	fmt.Printf("Dens.:\t %d ab./km²\n", capl[sig].densita)
	fmt.Printf("Alt. :\t %dm slm\n", capl[sig].altitudine)
}
