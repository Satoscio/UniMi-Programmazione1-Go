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

// newPunto crea un nuovo punto
func newPunto(x, y float64) Punto {
	var p3 Punto
	p3.x = x
	p3.y = y
	return p3
}

// specchiaPunto specchia rispetto all'asse Y un punto
func specchiaPunto(p *Punto) {
	if (*p).x == 0 {
		return
	}
	(*p).x *= -1
}
