package main

import "fmt"

// Punto è un punto nel piano cartesiano
type Punto struct {
	x, y float64
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	p3 := newPunto(a, b)
	fmt.Println(p3)
	specchiaPunto(&p3)
	fmt.Println(p3)
}

func newPunto(x, y float64) Punto {
	var p3 Punto
	p3.x = x
	p3.y = y
	return p3
}

func specchiaPunto(p *Punto) {
	(*p).x *= -1
}
