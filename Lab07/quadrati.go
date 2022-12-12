package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)


func isSquare(n int) bool {
    x := math.Sqrt(float64(n))
    z := math.Mod(x, float64(1))
    return z == 0
}

func main() {
    var nss = []string{}
    var nsn = []int{}

    nss = os.Args[1:]
    
    for _, n := range nss {
        i, _ := strconv.Atoi(n)
        if isSquare(i) {
            nsn = append(nsn, i)
        }
    }

    fmt.Println(nsn)
}