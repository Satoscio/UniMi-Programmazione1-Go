package main

import "fmt"

func main() {
	var gg1, gg2, st1, st2, e1, e2 int
	fmt.Print("Appuntamento 1 (gg, start, end): ")
	fmt.Scan(&gg1, &st1, &e1)
	fmt.Print("Appuntamento 2 (gg, start, end): ")
	fmt.Scan(&gg2, &st2, &e2)
	if gg1 != gg2 {
		fmt.Println("Non si sovrappongono")
	} else {
		if (st1 >= st2 && st1 <= e2) || (st2 >= st1 && st2 <= e1) {
			fmt.Println("Si sovrappongono")
		} else {
			fmt.Println("Non si sovrappongono")
		}
	}
}
