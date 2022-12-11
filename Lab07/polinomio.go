package main

import (
	"fmt"
	"math"
)

func main() {
	numeri := []float64{}
	var n, ris float64

	for {
		num, _ := fmt.Scan(&n)
		if num == 0 {
			break
		}
		numeri = append(numeri, n)
	}

	x := numeri[len(numeri)-1]
	numeri = numeri[:len(numeri)-1]

	for i, c := range numeri {
		switch i {
		case 0:
			ris += c
			fmt.Printf("%.1f + ", c)
		case 1:
			ris += c * x
			fmt.Printf("%.1fx + ", c)
		default:
			ris += c * math.Pow(x, float64(i))
			fmt.Printf("%.1fx^%d", c, i)
			if i != len(numeri)-1 {
				fmt.Print(" + ")
			}
		}
	}
	fmt.Printf("\nDa evaluare per x=%.0f\n", x)
	fmt.Println("Risulta:", ris)
}
