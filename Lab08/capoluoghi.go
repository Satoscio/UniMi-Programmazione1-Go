package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Capoluogo struct {
	nome, sigla, regione string
	popolazione, superficie, densita, altitudine int
}

func main() {
	capl := map[string]Capoluogo{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // serve a saltare la riga dell'intestazione

	for scanner.Scan() {
		var row Capoluogo
		scn := scanner.Text()
		all := strings.Split(scn, ",")
		als := all[:3] // parte in string
		aln := all[3:]
		nms := []int{} // parte in integer

		for _, i := range aln {
			j, _ := strconv.Atoi(i)
			nms = append(nms, j)
		}
	
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
	fmt.Printf("Nome: %s\n", capl[sig].nome)
	fmt.Printf("Sigla: %s\n", capl[sig].sigla)
	fmt.Printf("Reg.: %s\n", capl[sig].regione)
	fmt.Printf("Abit.: %d\n", capl[sig].popolazione)
	fmt.Printf("Area: %d\n", capl[sig].superficie)
	fmt.Printf("Dens.: %d ab/kmq\n", capl[sig].densita)
	fmt.Printf("Alt.: %d m slm\n", capl[sig].altitudine)
}