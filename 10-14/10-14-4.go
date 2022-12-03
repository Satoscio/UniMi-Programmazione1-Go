//  risolvere una equazione di secondo grado

package main

import (
    "fmt"
    "math"
    )

func main() {
    var a, b, c, x1, x2, delta float64

    fmt.Print("Inserisci i coefficienti di una equazione di secondo grado: ")
    fmt.Scan(&a, &b, &c)
    delta = b*b - (4 * a * c)
    if delta > 0 {
        x1 = (-b + math.Sqrt(delta)) / (2 * a)
        x2 = (-b - math.Sqrt(delta)) / (2 * a)
        fmt.Println("Primo risultato:", x1)
        fmt.Println("Secondo risultato:", x2)

    } else if delta == 0 {
        x1 = -b / (2 * a)
        _ = x2
        fmt.Println("Risultato:", x1)
    } else {
        fmt.Println("Equazione impossibile")
    }

}
