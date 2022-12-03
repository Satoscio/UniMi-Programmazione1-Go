// Leggo un intero e calcolare quante cifre ha

package main

import "fmt"

func main() {
    var x, n int
    fmt.Scan(&x)

    for x != 0 {
        x /= 10
        n++
    }

    fmt.Println(n)
}
