package main

import (
	"fmt"
	"math"
)


func isSquare(n int) bool {
    x := math.Sqrt(float64(n))
    z := math.Mod(x, float64(1))
    return !(z != 0)
}

func main() {
    var n int
    var ns = []int{}

    for {
        num, _ := fmt.Scan(&n)
        if num == 0 {
            break
        }
        if isSquare(n) {
            ns = append(ns, n)
        }
    }

    fmt.Println(ns)
}