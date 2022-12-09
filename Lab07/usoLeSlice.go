package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	x := []string{os.Args[1], os.Args[2], os.Args[3]}
	y := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		y = append(y, scanner.Text())
	}

	// .1
	z := append(x, y...)
	fmt.Println("x: ",x)
	fmt.Println("y: ",y)
	fmt.Println("01.", z)

	// .2
	sort.Strings(z)
	fmt.Println("02.", z)

	// .3
	z = z[:len(z)-1]
	fmt.Println("03.", z)

	
	// .4
	z = append(z[:2], z[4:]...)
	fmt.Println("04.", z)

	// .5
	a := []string{"a", "b", "c"}
	fmt.Println("05.", a)

	// .6
	z = append(z[:1], append(a, z[1:]...)...)
	fmt.Println("06.", z)

	// .7
	var s string
	fmt.Print("Nuova parola: ")
	fmt.Scan(&s)
	
	// per inserire un elemento x non slice in una slice, bisogna fare una specie
	// di casting []type{x}
	z = append(z[:2], append([]string{s}, z[2:]...)...)
	fmt.Println("07.", z)

	// .8
	var s1 string
	fmt.Print("Nuova parola: ")
	fmt.Scan(&s1)

	z = append(z, s1)
	fmt.Println("08.", z)
	
	// .9
	z = append(z, make([]string ,2)...)
	fmt.Println("09.", z)

	// .10
	z = append(z[:3], append(make([]string, 3), z[3:]...)...)
	fmt.Println("10.", z)

	// .11
	c := make([]string, len(z))
	copy(c, z)
	fmt.Println("11.", c)

	// .12
	c = c[:len(c)-1]
	fmt.Println("12.", z)
	fmt.Println("12.", c)
}