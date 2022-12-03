package main
import "fmt"
func main() {
    /*
        Scrivere un programma in Go che dato un intero n, stampi tutti gli
        interi pari minori o uguali a n
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
