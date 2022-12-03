package main
import "fmt"
func main() {
    /*
        Scrivere un programma in Go che prenda in input un numero intero n e
        stampi n volte un asterisco
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
