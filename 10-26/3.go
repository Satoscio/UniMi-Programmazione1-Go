package main

import (
    "fmt"
    "math"
)

func isPrime(n int) bool {
    for d := 2; d < n; d++ {
        if n % d == 0 {
            return false
        }
    }
    return true
}

func pi(x int) int {
    var c int
    for i := 1; i <= x; i++ {
        if isPrime(i) {
            c++
        }
    }
    return c
}

func main() {
    x := 1
    for {
        p := pi(x)
        pApprox := float64(x) / math.Log(float64(x))
        ratio := float64(p) / pApprox
        fmt.Println(x, "\t", p, "\t", pApprox, "\t", ratio)
    }
}
