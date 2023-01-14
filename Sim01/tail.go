package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	righe := []string{}
	n, _ := strconv.Atoi(os.Args[1])
	fn := os.Args[2]

	f, err := os.Open(fn)
	defer f.Close()
	if err != nil {
		fmt.Println("errore di lettura del file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		righe = append(righe, scanner.Text())
	}

	for i, r := range righe {
		if i > len(righe)-n-1 {
			fmt.Println(r)
		}
	}
}
