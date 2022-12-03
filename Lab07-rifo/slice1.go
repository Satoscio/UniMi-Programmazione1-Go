package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    parole := []string{}

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() && scanner.Text() != "stop" {
        parole = append(parole, scanner.Text())
    }

    for i := 0; i < len(parole); i++ {
        fmt.Println(parole[i])
    }

}
