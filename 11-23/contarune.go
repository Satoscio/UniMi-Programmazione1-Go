// incompleto

package main

import (
	"bufio"
	"fmt"
	"os"
)

func contaRune(s string) map[rune]int    {
    var r map[rune]int
    r = make(map[rune]int)

    for _, c := range s {
        r[c]++
    }

    return r
}

func printMap(m map[rune]int, totaleCar int) {
    for k, v := range m {
        perc := 100 * float64(v) / float64(totaleCar)
        fmt.Printf("%c: %8d %6.2f%%\t", k, v)
        for i := 0; i < int(perc*3); i++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    s := ""

    for scanner.Scan() {
        s += scanner.Text()
    }

    m := contaRune(os.Args[1])
    printMap(m, totaleCar)
}
