package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 int
	fmt.Print("P1: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("P2: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("P3: ")
	fmt.Scan(&x3, &y3)
	x1 = int(math.Min(float64(x1), float64(x2)))
	y1 = int(math.Min(float64(y1), float64(y2)))
	x2 = int(math.Max(float64(x1), float64(x2)))
	y2 = int(math.Max(float64(y1), float64(y2)))
	if x3 > x1 && x3 < x2 {
		if y3 > y1 && y3 < y2 {
			fmt.Println("P3 Interno")
		} else {
			fmt.Println("P3 Esterno")
		}
	} else if (x3 == x1 || x3 == x2) && (y3 > y1 && y3 < y2) {
		fmt.Println("P3 Perimetrale")
	} else if (y3 == y1 || y3 == y2) && (x3 > x1 && x3 < x2) {
		fmt.Println("P3 Perimetrale")
	} else {
		fmt.Println("P3 Esterno")
	}

}
