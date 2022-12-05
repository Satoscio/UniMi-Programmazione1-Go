package main

import "fmt"

const DIM = 5

func reverse(x *[DIM]int) {
	for i := 0; i < DIM/2; i++ {
		x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
	}
}

func switchFirstLast(x *[DIM]int) {
	x[0], x[len(x)-1] = x[len(x)-1], x[0]
}

func main() {
	var x = [DIM]int{5, 6, 9, 2, 0}
	fmt.Println("array:\t\t", x)
	reverse(&x)
	fmt.Println("reverse:\t", x)
	switchFirstLast(&x)
	fmt.Println("switchFirstLast:", x)
}
