package main

import (
    "fmt"
    "os"
    "bufio" // buffered input/output
)

func main() {
    parole := []string{}

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() && scanner.Text() != "stop" {
        parole = append(parole, scanner.Text())
    }

    for _, k := range parole {
        fmt.Println(k)
    }
}
