package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    parole := []string{}
    parolePari := []string{}

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() && scanner.Text() != "stop" {
        parole = append(parole, scanner.Text())
    }

    for i := 0; i < len(parole); i++ {
        if len(parole[i]) % 2 == 0 {
            parolePari = append(parolePari, parole[i])
        }
    }

    fmt.Println(parolePari)
}
