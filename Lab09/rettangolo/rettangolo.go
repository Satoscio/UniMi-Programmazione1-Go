package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rettangolo struct {
	base, altezza int
}

func (ret Rettangolo) String() (dis string) {
	for i := 0; i < ret.altezza; i++ {
		for i1 := 0; i1 < ret.base; i1++ {
			dis += "."
		}
		dis += "\n"
	}
	return
}

func main() {
	var ret Rettangolo
	x := os.Args[1:3]
	ret.base, _ = strconv.Atoi(x[0])
	ret.altezza, _ = strconv.Atoi(x[1])
	fmt.Println(ret.String())
}