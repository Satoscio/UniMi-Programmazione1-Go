package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main() {
	if len(os.Args) != 3 {
		fmt.Println("Inserire DUE nomi di file")
		os.Exit(1)
	}
	nf1 := os.Args[1]
	nf2 := os.Args[2]
	f1, e1 := os.Open(nf1)
	f2, e2 := os.Open(nf2)
	defer f1.Close()
	defer f2.Close()
	if e1 != nil || e2 != nil {
		fmt.Println("Errore di lettura di uno dei file")
		os.Exit(1)
	}
	sc1 := bufio.NewScanner(f1)
	sc2 := bufio.NewScanner(f2)

	sc1.Scan()
	sl1 := strings.Split(sc1.Text(), " ")
	sc2.Scan()
	sl2 := strings.Split(sc2.Text(), " ")
	sl3 := []string{}

	if len(sl1) > len(sl2) {
		sl3 = append(sl3, sl1[len(sl2):]...)
		sl1 = sl1[:len(sl2)]
	} else if len(sl1) < len(sl2) {
		sl3 = append(sl3, sl2[len(sl1):]...)
		sl2 = sl2[:len(sl1)]
	}

	for i, s := range sl1 {
		fmt.Println(s)
		fmt.Println(sl2[i])
	}

	if len(sl3) > 0 {
		for _, s := range sl3 {
			fmt.Println(s)
		}
	}
}