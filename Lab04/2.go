package main
import "fmt"

func main() {
    const DIM = 4
    var curr, prev int

    fmt.Scan(&curr)
    s := 0
    for i := 1; i < DIM; i++ {
        prev = curr
        fmt.Scan(&curr)
        d := curr - prev
        s += d
    }
    fmt.Println("risultato:", s)
}
