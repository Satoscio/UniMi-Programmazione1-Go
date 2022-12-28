package main

import "fmt"

func main() {
	var x, y float64
	fmt.Print("X Y: ")
	fmt.Scan(&x, &y)
	if x < 0 {
		if y < 0 {
			fmt.Println("III Quadrante")
		} else {
			fmt.Println("II Quadrante")
		}
	} else {
		if y < 0 {
			fmt.Println("IV Quadrante")
		} else {
			fmt.Println("I Quadrante")
		}
	}
}
